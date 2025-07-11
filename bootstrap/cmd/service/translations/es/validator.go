package es

import "github.com/gookit/validate"

const (
	Name          = "es"
	defaultErrMsg = "El campo {field} no pasó la validación"
)

// Register language data to validate.Validation
func Register(v *validate.Validation) {
	v.AddMessages(Data)
}

// RegisterGlobal register to validate global messages
func RegisterGlobal() {
	validate.AddGlobalMessages(Data)
}

// Data es language messages
var Data = map[string]string{
	"_": "{field}" + defaultErrMsg, // default message
	// builtin
	"_validate": "{field} no pasó la validación",      // default validate message
	"_filter":   "Los datos de {field} son inválidos", // data filter error
	// int value
	"min": "El valor mínimo de {field} es %v",
	"max": "El valor máximo de {field} es %v",
	// type check: int
	"isInt":  "El valor de {field} debe ser un número entero",
	"isInt1": "El valor de {field} debe ser un número entero y el valor mínimo es %d",     // has min check
	"isInt2": "El valor de {field} debe ser un número entero y estar en el rango %d - %d", // has min, max check
	"isInts": "El valor de {field} debe ser una lista de números enteros",
	"isUint": "El valor de {field} debe ser un número entero sin signo (>= 0)",
	// type check: string
	"isString":  "El valor de {field} debe ser una cadena de texto",
	"isString1": "El valor de {field} debe ser una cadena de texto y la longitud mínima es %d", // has min len check
	// length
	"minLength": "La longitud mínima de {field} es %d",
	"maxLength": "La longitud máxima de {field} es %d",
	// string length. calc rune
	"stringLength":  "La longitud de {field} debe estar en el rango %d - %d",
	"stringLength1": "La longitud mínima de {field} es %d",
	"stringLength2": "La longitud de {field} debe estar en el rango %d - %d",

	"isURL":     "{field} debe ser una dirección URL válida",
	"isFullURL": "{field} debe ser una dirección URL completa válida",
	"regexp":    "{field} debe coincidir con el patrón %s",

	"isFile":  "{field} debe ser un archivo subido",
	"isImage": "{field} debe ser un archivo de imagen subido",

	"enum":  "El valor de {field} debe estar en la enumeración %v",
	"range": "El valor de {field} debe estar en el rango %d - %d",
	// int compare
	"lt": "El valor de {field} debe ser menor que %v",
	"gt": "El valor de {field} debe ser mayor que %v",
	// required
	"required":           "{field} es obligatorio y no puede estar vacío",
	"requiredIf":         "{field} es obligatorio cuando {args0} está en {args1end}",
	"requiredUnless":     "El campo {field} es obligatorio a menos que {args0} esté en {args1end}",
	"requiredWith":       "El campo {field} es obligatorio cuando {values} está presente",
	"requiredWithAll":    "El campo {field} es obligatorio cuando {values} está presente",
	"requiredWithout":    "El campo {field} es obligatorio cuando {values} no está presente",
	"requiredWithoutAll": "El campo {field} es obligatorio cuando ninguno de {values} está presente",
	// field compare
	"eqField":  "El valor de {field} debe ser igual al campo %s",
	"neField":  "El valor de {field} no puede ser igual al campo %s",
	"ltField":  "El valor de {field} debe ser menor que el campo %s",
	"lteField": "El valor de {field} debe ser menor o igual que el campo %s",
	"gtField":  "El valor de {field} debe ser mayor que el campo %s",
	"gteField": "El valor de {field} debe ser mayor o igual que el campo %s",
	// data type
	"bool":    "El valor de {field} debe ser un booleano",
	"float":   "El valor de {field} debe ser un número decimal",
	"slice":   "El valor de {field} debe ser una lista",
	"map":     "El valor de {field} debe ser un mapa",
	"array":   "El valor de {field} debe ser un arreglo",
	"strings": "El valor de {field} debe ser una lista de cadenas de texto []string",
	"notIn":   "El valor de {field} no debe estar en la lista de enumeración dada %d",
	//
	"contains":    "El valor de {field} no contiene %s",
	"notContains": "El valor de {field} contiene %s",
	"startsWith":  "El valor de {field} no comienza con %s",
	"endsWith":    "El valor de {field} no termina con %s",
	"email":       "El valor de {field} es una dirección de correo electrónico inválida",
	"regex":       "El valor de {field} no pasa la verificación de expresión regular",
	"file":        "El valor de {field} debe ser un archivo",
	"image":       "El valor de {field} debe ser una imagen",
	// date
	"date":    "El valor de {field} debe ser una cadena de fecha",
	"gtDate":  "El valor de {field} debe ser posterior a %s",
	"ltDate":  "El valor de {field} debe ser anterior a %s",
	"gteDate": "El valor de {field} debe ser posterior o igual a %s",
	"lteDate": "El valor de {field} debe ser anterior o igual a %s",
	// check char
	"hasWhitespace":  "El valor de {field} debe contener espacios",
	"ascii":          "El valor de {field} debe ser una cadena ASCII",
	"alpha":          "El valor de {field} solo debe contener caracteres alfabéticos",
	"alphaNum":       "El valor de {field} solo debe contener caracteres alfabéticos y números",
	"alphaDash":      "El valor de {field} solo debe contener letras, números, guiones (-) y guiones bajos (_)",
	"multiByte":      "El valor de {field} debe ser una cadena multibyte",
	"base64":         "El valor de {field} debe ser una cadena base64",
	"dnsName":        "El valor de {field} debe ser una cadena DNS",
	"dataURI":        "El valor de {field} debe ser una cadena DataURL",
	"empty":          "El valor de {field} debe estar vacío",
	"hexColor":       "El valor de {field} debe ser una cadena de color en hexadecimal",
	"hexadecimal":    "El valor de {field} debe ser una cadena hexadecimal",
	"json":           "El valor de {field} debe ser una cadena JSON",
	"lat":            "El valor de {field} debe ser una coordenada de latitud",
	"lon":            "El valor de {field} debe ser una coordenada de longitud",
	"num":            "El valor de {field} debe ser una cadena numérica (>=0)",
	"mac":            "El valor de {field} debe ser una dirección MAC",
	"cnMobile":       "El valor de {field} debe ser una cadena de números de teléfono móvil chino de 11 dígitos",
	"printableASCII": "El valor de {field} debe ser una cadena ASCII imprimible",
	"rgbColor":       "El valor de {field} debe ser una cadena de color RGB",
	"fullURL":        "El valor de {field} debe ser una cadena URL completa",
	"full":           "El valor de {field} debe ser una cadena URL",
	"ip":             "El valor de {field} debe ser una cadena IP (v4 o v6)",
	"ipv4":           "El valor de {field} debe ser una cadena IPv4",
	"ipv6":           "El valor de {field} debe ser una cadena IPv6",
	"CIDR":           "El valor de {field} debe ser una cadena CIDR",
	"CIDRv4":         "El valor de {field} debe ser una cadena CIDRv4",
	"CIDRv6":         "El valor de {field} debe ser una cadena CIDRv6",
	"uuid":           "El valor de {field} debe ser una cadena UUID",
	"uuid3":          "El valor de {field} debe ser una cadena UUID3",
	"uuid4":          "El valor de {field} debe ser una cadena UUID4",
	"uuid5":          "El valor de {field} debe ser una cadena UUID5",
	"filePath":       "El valor de {field} debe ser una ruta de archivo existente",
	"unixPath":       "El valor de {field} debe ser una cadena de ruta unix",
	"winPath":        "El valor de {field} debe ser una cadena de ruta de Windows",
	"isbn10":         "El valor de {field} debe ser una cadena isbn10",
	"isbn13":         "El valor de {field} debe ser una cadena isbn13",
	"required_uuid":  "{field} es obligatorio y no puede estar vacío",
}
