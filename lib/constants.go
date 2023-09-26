package lib

const (
	XMLNS_PREFIX = "xmlns"
	SML_PREFIX   = ""
	GML_PREFIX   = "g"
	SWE_PREFIX   = "s"
	XLINK_PREFIX = "x"

	OGC_NS_PREFIX = "http://www.opengis.net/"
	SML_NS        = "sensorml/2.0"
	SWE_NS        = "swe/2.0"
	GML_NS        = "gml/3.2"
	XLINK_NS      = "http://www.w3.org/1999/xlink"

	TIME     = "Time"
	QUANTITY = "Quantity"
	COUNT    = "Count"
	CATEGORY = "Category"
	BOOLEAN  = "Boolean"
	TEXT     = "Text"

	ATT_ID         = "id"
	ATT_NAME       = "name"
	ATT_DEFINITION = "definition"
	ATT_REFRAME    = "referenceFrame"
	ATT_AXISID     = "axisID"

	ELT_COMPONENT   = "PhysicalComponent"
	ELT_SYSTEM      = "PhysicalSystem"
	ELT_IDENTIFIER  = "identifier"
	ELT_DESCRIPTION = "description"
	ELT_OUTPUTS     = "outputs"
	ELT_OUTPUTLIST  = "OutputList"
	ELT_OUTPUT      = "output"
	ELT_DATARECORD  = "DataRecord"
	ELT_VECTOR      = "Vector"
	ELT_FIELD       = "field"
	ELT_COORDINATE  = "coordinate"
	ELT_LABEL       = "label"
	ELT_UOM         = "uom"
	ATT_HREF        = "href"
	ATT_CODE        = "code"

	HTTP_PREFIX       = "http://"
	OGC_DEF_PREFIX    = "http://www.opengis.net/def/"
	DEF_SAMPLING_TIME = "property/OGC/0/SamplingTime"
	DEF_ISO8601       = "uom/ISO-8601/0/Gregorian"
	DEF_UTC           = "trs/BIPM/0/UTC"
	DEF_EPSG4326      = "crs/EPSG/0/4326"
	DEF_EPSG4979      = "crs/EPSG/0/4979"
)