package fakesApppointmentsRepository

import (
	"ismaeldf/golang-gobarber/modules/appointments/infra/gorm/entities"
	"time"
)

type FakeAppointmentsRepository struct {
	appointments []entities.Appointment
}

func (r *FakeAppointmentsRepository) FindByDate(date time.Time) entities.Appointment {
	var appointment entities.Appointment

	for _, a := range r.appointments{
		if a.Date == date{
			return a
		}
	}

	return appointment
}

func (r *FakeAppointmentsRepository) Create(data entities.Appointment) (*entities.Appointment, error) {
	appointment, err := entities.NewAppointment(data.ProviderId, data.Date)
	if err != nil{
		return nil, err
	}

	r.appointments = append(r.appointments, *appointment)

	return appointment, nil
}

func (r *FakeAppointmentsRepository) All() []entities.Appointment {
	return r.appointments
}
