package main

import (
	"errors"
	"fmt"

	"lesson5/student/practice"
)

func main() {
	successCase()
	emptyMessageCase()
	missingConfigCase()
	brokenChannelCase()
}

// successCase: все каналы настроены, сообщение доставлено в каждый.
func successCase() {
	fmt.Println("== Случай 1: успешная рассылка ==")
	email := &practice.EmailNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "email"},
		Address:      "student@example.com",
	}
	push := &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "push"},
		DeviceID:     "phone-1",
	}

	err := practice.Dispatch([]practice.Notifier{email, push}, "Заказ готов")
	fmt.Println("ошибка:", err)
	fmt.Println("email история:", email.Messages)
	fmt.Println("push история:", push.Messages)
	fmt.Println()
}

// emptyMessageCase: пустой текст распознаётся через errors.Is как ErrEmptyMessage.
func emptyMessageCase() {
	fmt.Println("== Случай 2: пустой текст ==")
	email := &practice.EmailNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "email"},
		Address:      "student@example.com",
	}
	push := &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "push"},
		DeviceID:     "phone-1",
	}

	err := practice.Dispatch([]practice.Notifier{email, push}, "")
	if errors.Is(err, practice.ErrEmptyMessage) {
		fmt.Println("распознан ErrEmptyMessage:", err)
	} else {
		fmt.Println("неожиданная ошибка:", err)
	}
	fmt.Println("email история:", email.Messages)
	fmt.Println("push история:", push.Messages)
	fmt.Println()
}

// missingConfigCase: ошибка настройки извлекается через errors.As как *SendError.
func missingConfigCase() {
	fmt.Println("== Случай 3: нет настройки у push ==")
	push := &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "push"},
		DeviceID:     "",
	}

	err := practice.Dispatch([]practice.Notifier{push}, "Привет")
	var sendErr *practice.SendError
	if errors.As(err, &sendErr) {
		fmt.Printf("канал: %s, причина: %s\n", sendErr.Channel, sendErr.Reason)
	} else {
		fmt.Println("неожиданная ошибка:", err)
	}
	fmt.Println("push история:", push.Messages)
	fmt.Println()
}

// brokenChannelCase: паника одного канала не мешает следующему; SafeDispatch
// возвращает ошибку, которую распознаёт errors.Is через ErrNotifierPanic.
func brokenChannelCase() {
	fmt.Println("== Случай 4: сломанный канал и продолжение ==")
	push := &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "push"},
		DeviceID:     "phone-1",
	}
	broken := &practice.BrokenNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "broken"},
	}

	failures := practice.SafeDispatch([]practice.Notifier{broken, push}, "Заказ готов")
	if len(failures) == 0 {
		fmt.Println("ошибок нет")
	} else {
		for i, err := range failures {
			if errors.Is(err, practice.ErrNotifierPanic) {
				fmt.Printf("ошибка %d: распознан ErrNotifierPanic: %v\n", i+1, err)
			} else {
				fmt.Printf("ошибка %d: %v\n", i+1, err)
			}
		}
	}
	fmt.Println("push история:", push.Messages)
	fmt.Println()
}
