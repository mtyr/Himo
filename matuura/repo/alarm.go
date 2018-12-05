package repo

import "github.com/mtyr/Himo/matuura/model"

//var alarms []model.Alarm

var alarms []string

/*
func SetAlarm(alarm model.Alarm) {
	alarms = append(alarms, alarm)
}*/

func SetAlarm(alarm) {
	alarms = append(alarms, alarm)
}

func GetAllAlarm() []model.Alarm {
	return alarms
}

func FindBySeetIDAlarm(seet_id string) *model.Alarm {
	for _, alarm := range alarms {
		if alarm.SeetID == seet_id {
			return &alarm
		}
	}
	return nil
}
