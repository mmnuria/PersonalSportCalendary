package healthScheduler

type Empleado struct {
    Nombre       string
    TipoContrato Contrato
    Turno        Turno
    Area         Area
    Disponibilidad []string
}

// Función para verificar si un empleado tiene el día libre
func (e *Empleado) EsDiaLibre(dia string) bool {
    for _, diaLibre := range e.Disponibilidad {
        if diaLibre == dia {
            return true
        }
    }
    return false
}

func NuevoEmpleado(nombre, tipoContrato Contrato, turno Turno, area Area, disponibilidad []string) *Empleado {
    return &Empleado{
        Nombre:       nombre,
        TipoContrato: tipoContrato,
        Turno:        turno,
        Area:         area,
        Disponibilidad:   disponibilidad,
    }
}

