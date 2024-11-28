package PersonalSportCalendary

import "fmt"

type Rutina struct {
	Nombre     string      //Nombre de la rutina
	Tiempo     int         //Duración en minutos
	Ejercicios []Ejercicio //Instrucciones detalladas sobre los ejercicios que componen la rutina y sus repeticiones
	Materiales []Material  //Materiales necesarios para realizar la rutina
}

func NewRutina(nombre string, tiempo int, ejercicios []Ejercicio, materiales []Material) Rutina {
	return Rutina{
		Nombre:     nombre,
		Tiempo:     tiempo,
		Ejercicios: ejercicios,
		Materiales: materiales,
	}
}

type Material string //enum de materiales
const (
	Peso          = "Peso"
	BandaElastica = "Banda Elastica"
	Cuerda        = "Cuerda"
	Esterilla     = "Esterilla"
	Pelota        = "Pelota"
)

func NewMaterial(material string) (Material, error) {
	switch Material(material) {
	case Peso, BandaElastica, Cuerda, Esterilla, Pelota:
		return Material(material), nil
	default:
		return "", fmt.Errorf("Material inválido: %v", material)
	}
}
