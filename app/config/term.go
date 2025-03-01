package config

const termYearKey = "termYearKey"
const termKey = "termKey"
const termStartDate = "termStartDate"
const scoreTermYearKey = "scoreTermYearKey"
const scoreTermKey = "scoreTermKey"

func SetSystemInfo(yearValue, termValue, termStartDateValue, scoreYearValue, scoreTermValue, jpgUrlValue, fileUrlValue, registerTipsValue, schoolBusUrlValue string) error {
	err := setConfig(termYearKey, yearValue)
	if err != nil {
		return err
	}
	err = setConfig(termKey, termValue)
	if err != nil {
		return err
	}
	err = setConfig(termStartDate, termStartDateValue)
	if err != nil {
		return err
	}
	err = setConfig(scoreTermYearKey, scoreYearValue)
	if err != nil {
		return err
	}
	err = setConfig(scoreTermKey, scoreTermValue)
	if err != nil {
		return err
	}
	err = setConfig(schoolBusUrlKey, schoolBusUrlValue)
	if err != nil {
		return err
	}
	err = setConfig(webpUrlKey, jpgUrlValue)
	if err != nil {
		return err
	}
	err = setConfig(fileUrlKey, fileUrlValue)
	if err != nil {
		return err
	}
	err = setConfig(registerTipsKey, registerTipsValue)
	if err != nil {
		return err
	}
	return err
}

func GetTermInfo() (string, string, string, string, string) {
	return getConfig(termYearKey), getConfig(termKey), getConfig(termStartDate), getConfig(scoreTermYearKey), getConfig(scoreTermKey)
}

func IsSetTermInfo() bool {
	return checkConfig(termYearKey) && checkConfig(termKey) && checkConfig(termStartDate)
}

func DelTermInfo() []error {
	var result []error
	errTermYear := delConfig(termYearKey)
	errTerm := delConfig(termKey)
	errStartDate := delConfig(termStartDate)
	return append(result, errTermYear, errTerm, errStartDate)
}
