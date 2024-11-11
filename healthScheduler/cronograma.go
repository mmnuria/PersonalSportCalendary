package healthScheduler

type Asignacion struct {
    Empleado Empleado
    Area     Area
    Turno    Turno
    EsFlexible  bool
}

type CronogramaDiario struct {
    Fecha       string
    Asignaciones []Asignacion
}

func NuevaAsignacion(empleado Empleado, area Area, turno Turno, esFlexible bool) *Cronograma {
    return &Asignacion{
        Empleado: empleado,
        Area:     area,
        Turno:    turno,
        EsFlexible: esFlexible,
    }
}

func NuevoCronogramaDiario(fecha string, asignaciones []Asignacion) CronogramaDiario {
	return CronogramaDiario{
		Fecha:       fecha,
		Asignaciones: asignaciones,
	}
}
