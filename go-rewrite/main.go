package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"os"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type CarInClass struct {
	Car_ID   int    `json:"car_id"`
	Car_Name string `json:"car_name"`
}

type CarClass struct {
	Car_Class_ID  int          `json:"car_class_id"`
	Cars_In_Class []CarInClass `json:"cars_in_class"`
	Name          string       `json:"name"`
	Short_Name    string       `json:"short_name"`
}

type License struct {
	Parent_ID         int    `json:"parent_id"`
	License_Group     int    `json:"license_group"`
	Min_License_Level int    `json:"min_license_level"`
	Max_License_Level int    `json:"max_license_level"`
	Group_Name        string `json:"group_name"`
}

type RaceSummary struct {
	Subsession_ID           int    `json:"subsession_id"`
	Average_Lap             int    `json:"average_lap"`
	Laps_Complete           int    `json:"laps_complete"`
	Num_Cautions            int    `json:"num_cautions"`
	Num_Caution_Laps        int    `json:"num_caution_laps"`
	Num_Lead_Changes        int    `json:"num_lead_changes"`
	Field_Strength          int    `json:"field_strength"`
	Num_Opt_Laps            int    `json:"num_opt_laps"`
	Has_Opt_Path            bool   `json:"has_opt_path"`
	Special_Event_Type      int    `json:"special_event_type"`
	Special_Event_Type_Text string `json:"special_event_type_text"`
}

type Track struct {
	Track_ID      int    `json:"track_id"`
	Track_Name    string `json:"track_name"`
	Config_Name   string `json:"config_name"`
	Category_ID   int    `json:"category_id"`
	Category      string `json:"category"`
	Race_Week_Num int    `json:"race_week_num"`
}

type SessionResult struct {
	Cust_ID                    int     `json:"cust_id"`
	Display_Name               string  `json:"display_name"`
	Finish_Position            int     `json:"finish_position"`
	Finish_Position_In_Class   int     `json:"finish_position_in_class"`
	Laps_Lead                  int     `json:"laps_lead"`
	Laps_Complete              int     `json:"laps_complete"`
	Opt_Laps_Complete          int     `json:"opt_laps_complete"`
	Interval                   int     `json:"interval"`
	Class_Interval             int     `json:"class_interval"`
	Average_Lap                int     `json:"average_lap"`
	Best_Lap_Num               int     `json:"best_lap_num"`
	Best_Lap_Time              int     `json:"best_lap_time"`
	Best_Nlaps_Num             int     `json:"best_nlaps_num"`
	Best_Nlaps_Time            int     `json:"best_nlaps_time"`
	Best_Qual_Lap_At           string  `json:"best_qual_lap_at"`
	Best_Qual_Lap_Num          int     `json:"best_qual_lap_num"`
	Best_Qual_Lap_Time         int     `json:"best_qual_lap_time"`
	Reason_Out_ID              int     `json:"reason_out_id"`
	Reason_Out                 string  `json:"reason_out"`
	Champ_Points               int     `json:"champ_points"`
	Drop_Race                  bool    `json:"drop_race"`
	Position                   int     `json:"position"`
	Qual_Lap_Time              int     `json:"qual_lap_time"`
	Starting_Position          int     `json:"starting_position"`
	Starting_Position_In_Class int     `json:"starting_position_in_class"`
	Car_Class_ID               int     `json:"car_class_id"`
	Car_Class_Name             string  `json:"car_class_name"`
	Car_Class_Short_Name       string  `json:"car_class_short_name"`
	Division                   int     `json:"division"`
	Division_Name              string  `json:"division_name"`
	Old_License_Level          int     `json:"old_license_level"`
	Old_Sub_Level              int     `json:"old_sub_level"`
	Old_CPI                    float64 `json:"old_cpi"`
	Oldi_Rating                int     `json:"oldi_rating"`
	Old_TTrating               int     `json:"old_ttrating"`
	New_License_Level          int     `json:"new_license_level"`
	New_Sub_Level              int     `json:"new_sub_level"`
	New_CPI                    float64 `json:"new_cpi"`
	Newi_Rating                int     `json:"newi_rating"`
	New_TTrating               int     `json:"new_ttrating"`
	Multiplier                 int     `json:"multiplier"`
	License_Change_Oval        int     `json:"license_change_oval"`
	License_Change_Road        int     `json:"license_change_road"`
	Incidents                  int     `json:"incidents"`
	Max_PCT_Fuel_Fill          int     `json:"max_pct_fuel_fill"`
	Weight_Penalty_Kg          int     `json:"weight_penalty_kg"`
	League_Points              int     `json:"league_points"`
	League_Agg_Points          int     `json:"league_agg_points"`
	Car_ID                     int     `json:"car_id"`
	Car_Name                   string  `json:"car_name"`
	Aggregate_Champ_Points     int     `json:"aggregate_champ_points"`
}

type Session struct {
	Simsession_Name      string          `json:"simsession_name"`
	Simsession_Number    int             `json:"simsession_number"`
	Simsession_Type      int             `json:"simsession_type"`
	Simsession_Type_Name string          `json:"simsession_type_name"`
	Simsession_Subtype   int             `json:"simsession_subtype"`
	Results              []SessionResult `json:"results"`
}

type Subsession struct {
	Allow_Licenses          []License   `json:"allowed_licenses"`
	Car_Classes             []CarClass  `json:"car_classes"`
	Corners_Per_Lap         int         `json:"corners_per_lap"`
	End_Time                string      `json:"end_time"`
	Event_Average_Lap       int         `json:"event_average_lap"`
	Event_Laps_Complete     int         `json:"event_laps_complete"`
	Event_Strength_Of_Field int         `json:"event_strength_of_field"`
	License_Category        string      `json:"license_category"`
	License_Category_ID     int         `json:"license_category_id"`
	Max_Weeks               int         `json:"max_weeks"`
	Num_Caution_Laps        int         `json:"num_caution_Laps"`
	Num_Cautions            int         `json:"num_cautions"`
	Num_Lead_Changes        int         `json:"num_lead_changes"`
	Race_Summary            RaceSummary `json:"race_summary"`
	Race_Week_Num           int         `json:"race_week_num"`
	Season_ID               int         `json:"season_id"`
	Season_Name             string      `json:"season_name"`
	Season_Quarter          int         `json:"season_quarter"`
	Season_Short_Name       string      `json:"season_short_name"`
	Season_Year             int         `json:"season_year"`
	Series_ID               int         `json:"series_id"`
	Series_Logo             string      `json:"series_logo"`
	Series_Name             string      `json:"series_name"`
	Series_Short_Name       string      `json:"series_short_name"`
	Session_ID              int         `json:"session_id"`
	Session_Results         []Session   `json:"session_results"`
	Start_Time              string      `json:"start_time"`
	Subsession_ID           int         `json:"subsession_id"`
	Track                   Track       `json:"track"`
}

type HandleResultsOutput struct {
	Keys            []string
	Handled_Results []SessionResult
}

func convert_lap_time(lap_time int) string {
	ingested_lap := strconv.FormatFloat(float64(lap_time), 'f', 4, 64)
	with_decimal := ""
	if lap_time > 0 {
		with_decimal = strconv.FormatFloat(float64(lap_time)/10000.0, 'f', 4, 64)
	} else {
		with_decimal = ingested_lap
	}
	split_results := strings.Split(with_decimal, ".")
	if len(split_results[0]) > 2 {
		matcher, _ := regexp.Compile("[0-9]{2}$")
		match := matcher.FindString(split_results[0])
		strings.Replace(with_decimal, match, fmt.Sprintf(".%s", match), -1)
	}
	numerical_value, _ := strconv.ParseFloat(with_decimal, 64)
	// If the lap is a minute or longer
	if numerical_value >= 60.0 {
		minutes := math.Floor(numerical_value / 60.0)
		seconds := math.Floor(numerical_value - (minutes * 60.0))
		split_length := len(strings.Split(with_decimal, "."))
		subseconds := strings.Split(with_decimal, ".")[split_length-1]
		seconds_string := ""
		if seconds >= 10 {
			seconds_string = strconv.FormatFloat(float64(seconds), 'f', 0, 64)
		} else {
			seconds_string = fmt.Sprintf("0%s", strconv.FormatFloat(float64(seconds), 'f', 0, 64))
		}
		return fmt.Sprintf("%s:%s.%s", strconv.FormatFloat(float64(minutes), 'f', 0, 64), seconds_string, subseconds)
	}
	return strconv.FormatFloat(numerical_value, 'f', 3, 64)
}

func main() {
	raw_subsessions_input, err := os.ReadFile("./1-subsessions-output.json")
	if err != nil {
		fmt.Println(1, err)
	}
	var subsessions []Subsession
	err = json.Unmarshal(raw_subsessions_input, &subsessions)
	if err != nil {
		fmt.Println(2, err)
	}
	function_map := template.FuncMap{
		"bump": func(i int) int {
			return i + 1
		},
		"convert_time": func(timestamp string) string {
			// Note: This can be changed to use any timezone you so choose
			timezone, _ := time.LoadLocation("America/Chicago")
			converted_time, _ := time.ParseInLocation("2006-01-02T15:04:05Z", timestamp, time.UTC)
			const layout = "1/2/2006 3:04pm"
			return converted_time.In(timezone).Format(layout)
		},
		"convert_lap_time": func(lap_time int) string {
			return convert_lap_time(lap_time)
		},
		"filter_allowed_licenses": func(licenses []License) []string {
			filtered_allowed_licenses := []string{}
			for index, license := range licenses {
				if index < 2 {
					if index == 0 {
						filtered_allowed_licenses = append(filtered_allowed_licenses, fmt.Sprintf("%s %s.00", license.Group_Name, strconv.Itoa(license.Min_License_Level/license.License_Group)))
					} else {
						filtered_allowed_licenses = append(filtered_allowed_licenses, fmt.Sprintf("%s and above", license.Group_Name))
					}
				}
			}
			return filtered_allowed_licenses
		},
		"find_session": func(sessions []Session, session_name string) Session {
			session_to_return := Session{}
			if len(sessions) >= 4 {
				if session_name == "RACE" {
					session_to_return = sessions[len(sessions)-1]
					session_to_return.Simsession_Name = "RACE"
					return session_to_return
				} else if session_name == "HEAT" {
					session_to_return = sessions[len(sessions)-2]
					session_to_return.Simsession_Name = "HEAT"
					return session_to_return
				}
			}
			for _, session := range sessions {
				if session.Simsession_Name == session_name {
					session_to_return = session
				}
			}
			return session_to_return
		},
		"retrieve_keys_to_display": func(car_classes []CarClass, session_name string) []string {
			// Admittedly this whole function is kinda lazy but meh
			if len(car_classes) > 1 && session_name == "RACE" {
				return []string{
					"cust_id",
					"display_name",
					"finish_position",
					"finish_position_in_class",
					"laps_lead",
					"laps_complete",
					"interval",
					"class_interval",
					"average_lap",
					"best_lap_num",
					"best_lap_time",
					"reason_out",
					"champ_points",
					"drop_race",
					"starting_position",
					"starting_position_in_class",
					"car_class_short_name",
					"division_name",
					"old_license_level",
					"old_cpi",
					"oldi_rating",
					"new_license_level",
					"new_cpi",
					"newi_rating",
					"incidents",
					"car_name",
					"aggregate_champ_points",
				}
			} else if len(car_classes) > 1 {
				return []string{
					"cust_id",
					"display_name",
					"finish_position",
					"finish_position_in_class",
					"laps_complete",
					"interval",
					"class_interval",
					"average_lap",
					"best_lap_num",
					"best_lap_time",
					"car_class_short_name",
					"division_name",
					"incidents",
					"car_name",
				}
			} else if session_name == "RACE" {
				return []string{
					"cust_id",
					"display_name",
					"finish_position",
					"laps_lead",
					"laps_complete",
					"interval",
					"average_lap",
					"best_lap_num",
					"best_lap_time",
					"reason_out",
					"champ_points",
					"drop_race",
					"starting_position",
					"division_name",
					"old_license_level",
					"old_cpi",
					"oldi_rating",
					"new_license_level",
					"new_cpi",
					"newi_rating",
					"incidents",
					"car_name",
					"aggregate_champ_points",
				}
			}
			return []string{
				"cust_id",
				"display_name",
				"finish_position",
				"laps_complete",
				"interval",
				"average_lap",
				"best_lap_num",
				"best_lap_time",
				"division_name",
				"incidents",
				"car_name",
			}
		},
		"retrieve_key_labels": func(keys []string) []string {
			key_map := make(map[string]string)
			key_map["display_name"] = "Name"
			key_map["finish_position"] = "Finish"
			key_map["finish_position_in_class"] = "Finish (In Class)"
			key_map["laps_lead"] = "Laps Lead"
			key_map["laps_complete"] = "Laps Completed"
			key_map["interval"] = "Interval to P1"
			key_map["class_interval"] = "Interval to P1 (In Class)"
			key_map["average_lap"] = "AVG Laptime"
			key_map["best_lap_num"] = "Fastest Lap"
			key_map["best_lap_time"] = "Fastest Lap Time"
			key_map["reason_out"] = "Status"
			key_map["champ_points"] = "Points"
			key_map["drop_race"] = "Drop Race"
			key_map["starting_position"] = "Starting Position"
			key_map["starting_position_in_class"] = "Starting Position (In Class)"
			key_map["car_class_short_name"] = "Car Class"
			key_map["division_name"] = "Division"
			key_map["old_license_level"] = "Old License Level"
			key_map["old_cpi"] = "Old Corners Per Incident"
			key_map["oldi_rating"] = "Old iRating"
			key_map["new_license_level"] = "New License Level"
			key_map["new_cpi"] = "New Corners Per Incident"
			key_map["newi_rating"] = "New iRating"
			key_map["incidents"] = "Incidents"
			key_map["car_name"] = "Car"
			key_map["aggregate_champ_points"] = "Best Points For Week"
			key_map["license_category"] = "License Category"
			key_map["year"] = "Year"
			key_map["season_number_of_year"] = "Season Number"
			key_map["car_class_id"] = "Car Class ID"
			key_map["division"] = "Division"
			key_map["division_rank"] = "Rank (Division)"
			key_map["overall_rank"] = "Rank (Overall)"
			key_map["season_id"] = "Season ID"
			key_map["season_name"] = "Season Name"
			key_map["season_summary"] = "Season Summary"
			key_map["weeks_counted"] = "Participation Weeks"
			key_map["starts"] = "Starts"
			key_map["wins"] = "Wins"
			key_map["top5"] = "Top 5's"
			key_map["top25_percent"] = "Top 25% Finishes"
			key_map["poles"] = "Poles"
			key_map["avg_start_position"] = "Average Starting Position"
			key_map["avg_finish_position"] = "Average Finish Position"
			key_map["avg_field_size"] = "Average Field Size"
			key_map["laps"] = "Laps Completed"
			key_map["laps_led"] = "Laps Led"
			key_map["points"] = "Points"
			key_map["category"] = "License Category"
			key_map["safety_rating"] = "Safety Rating"
			key_map["irating"] = "iRating"
			key_map["group_name"] = "License Group Name"
			key_map["race_week_num"] = "Race Week Number"
			key_map["num_lead_changes"] = "Number of Lead Changes"
			key_map["num_cautions"] = "Number of Cautions"
			key_map["num_caution_laps"] = "Number of Caution Laps"
			key_map["event_strength_of_field"] = "Strength of Field"
			key_map["event_laps_complete"] = "Total Laps"
			key_map["event_average_lap"] = "Average Lap Time"
			output_labels := []string{}
			for _, key := range keys {
				output_labels = append(output_labels, key_map[strings.ToLower(key)])
			}
			return output_labels
		},
		"handle_results": func(keys_to_display []string, results []SessionResult) HandleResultsOutput {
			unique_key_map := make(map[string]string)
			handled_results := []SessionResult{}
			for _, result := range results {
				reflected_result := reflect.ValueOf(&result).Elem()
				fields := reflect.VisibleFields(reflect.TypeOf(result))
				for _, field := range fields {
					if slices.Contains(keys_to_display, strings.ToLower(field.Name)) {
						unique_key_map[strings.ToLower(field.Name)] = field.Name
						// Being lazy because types are cumbersome sometimes
						if strings.Contains(field.Name, "Position") {
							field_value := reflected_result.FieldByName(field.Name)
							field_value.SetInt(field_value.Int() + 1)
						}
					}
				}
				handled_results = append(handled_results, result)
			}
			keys := make([]string, 0, len(unique_key_map))
			for _, k := range keys_to_display {
				if unique_key_map[k] != "" {
					keys = append(keys, unique_key_map[k])
				}
			}
			return HandleResultsOutput{keys, handled_results}
		},
		"return_value_at_key": func(result SessionResult, key string) any {
			reflected_result := reflect.ValueOf(&result).Elem()
			value := reflected_result.FieldByName(key)
			if slices.Contains([]string{"Interval", "Class_Interval", "Average_Lap", "Best_Lap_Time"}, key) {
				return convert_lap_time(int(value.Int()))
			}
			return value
		},
		"return_key_value_at_index": func(keys []string, index int) string {
			return keys[index]
		},
	}
	template_file := "subsessions.html"
	html_template, err := template.New(template_file).Funcs(function_map).ParseFiles(template_file)
	for subsession_index, subsession := range subsessions {
		fmt.Println(fmt.Sprintf("Creating %s file of %s", strconv.Itoa(subsession_index), strconv.Itoa(len(subsessions))))
		os.MkdirAll(fmt.Sprintf("./subsession/%s", strconv.Itoa(subsession.Subsession_ID)), os.ModePerm)
		file, err := os.Create(fmt.Sprintf("./subsession/%s/index.html", strconv.Itoa(subsession.Subsession_ID)))
		if err != nil {
			fmt.Println(3, err)
		}
		// fmt.Printf("%+v\n", subsession)
		err = html_template.Execute(file, subsession)
		if err != nil {
			fmt.Println(4, err)
		}
		file.Close()
	}
}
