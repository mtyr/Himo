package repo

import "../model"

var alarms []model.Alarm

func SetAlarm(alarm model.Alarm) {
	alarms = append(alarms, alarm)
}

func GetAllAlarm() []model.Alarm {
	return alarms
}

func SetSeat() []model.Seat {
	return seats
}

func FindBySeetIDAlarm(seet_id string) *model.Alarm {
	for _, alarm := range alarms {
		if alarm.SeetID == seet_id {
			return &alarm
		}
	}
	return nil
}
