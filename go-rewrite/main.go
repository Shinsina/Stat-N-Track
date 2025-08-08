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
	"sort"
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

type Standing struct {
	Car_Class_ID       int              `json:"car_class_id"`
	Division           int              `json:"division"`
	Overall_Rank       int              `json:"overall_rank"`
	Season_Driver_Data SeasonDriverData `json:"season_driver_data"`
	Season_ID          int              `json:"season_id"`
	Season_Name        string           `json:"season_name"`
	Division_Rank      int              `json:"division_rank"`
}

type SeasonDriverData struct {
	Rank                int    `json:"rank"`
	Cust_ID             int    `json:"cust_id"`
	Display_Name        string `json:"display_name"`
	Weeks_Counted       int    `json:"weeks_counted"`
	Starts              int    `json:"starts"`
	Wins                int    `json:"wins"`
	Top5                int    `json:"top5"`
	Top25_Percent       int    `json:"top25_percent"`
	Poles               int    `json:"poles"`
	Avg_Start_Position  int    `json:"avg_start_position"`
	Avg_Finish_Position int    `json:"avg_finish_position"`
	Avg_Field_Size      int    `json:"avg_field_size"`
	Laps                int    `json:"laps"`
	Laps_Led            int    `json:"laps_led"`
	Incidents           int    `json:"incidents"`
	Points              int    `json:"points"`
}

type StandingWithSubsessions struct {
	Standing    Standing
	Subsessions []Subsession
}

type ConsolidatedSubsessionResult struct {
	Subsession_ID              int
	Track                      string
	Corners_Per_Lap            int
	License                    string
	License_Category           string
	Finish_Position            int
	Finish_Position_In_Class   int
	Laps_Lead                  int
	Laps_Complete              int
	Average_Lap                int
	Best_Lap_Time              int
	Reason_Out                 string
	Champ_Points               int
	Starting_Position          int
	Starting_Position_In_Class int
	Car_Class_Short_Name       string
	Division_Name              string
	New_License_Level          int
	New_CPI                    float64
	Newi_Rating                int
	Incidents                  int
	Car_Name                   string
	Aggregate_Champ_Points     int
}

type ConsolidatedHeadToHeadSubsessionResult struct {
	Subsession_ID           int
	Track                   string
	Corners_Per_Lap         int
	License                 string
	License_Category        string
	Event_Average_Lap       int
	Event_Laps_Complete     int
	Event_Strength_Of_Field int
	Num_Caution_Laps        int
	Num_Cautions            int
	Num_Lead_Changes        int
	Race_Week_Num           int
	Season_Name             string
}

type HandleConsolidatedHeadToHeadResultsOutput struct {
	Keys                         []string
	Handled_Consolidated_Results []ConsolidatedHeadToHeadSubsessionResult
}

type HeadToHeadWinsOutput struct {
	Key       string
	Win_Count int
}

type HandleConsolidatedResultsOutput struct {
	Keys                         []string
	Handled_Consolidated_Results []ConsolidatedSubsessionResult
}

type RaceWeek struct {
	Race_Week_Num int   `json:"race_week_num"`
	Track         Track `json:"track"`
}

type Season struct {
	Season_ID  int        `json:"season_id"`
	Race_Weeks []RaceWeek `json:"race_weeks"`
}

type Series struct {
	Seasons []Season `json:"seasons"`
}

type PastSeason struct {
	Series_ID int    `json:"series_id"`
	Series    Series `json:"series"`
}

type SummaryStats struct {
	Avg_Start  string
	Avg_Finish string
	Wins       int
	Incidents  int
}

type SubsessionListSheetData struct {
	Description          string
	Subsessions          []Subsession
	Stylesheet_Path      string
	Table_Component_Path string
}

type ConsolidatedStandingResult struct {
	Season_ID           int
	Year                int
	Season_Number       int
	Car_Class_ID        int
	Season_Summary      int
	Division            int
	Division_Rank       int
	Overall_Rank        int
	Points              int
	Season_Name         string
	Weeks_Counted       int
	Starts              int
	Wins                int
	Top5                int
	Top25_Percent       int
	Poles               int
	Avg_Start_Position  int
	Avg_Finish_Position int
	Avg_Field_Size      int
	Laps                int
	Laps_Led            int
	Incidents           int
	Cust_ID             int
}

type StandingListSheetData struct {
	Description          string
	Standings            []Standing
	Stylesheet_Path      string
	Table_Component_Path string
}

type HandleConsolidatedStandingResultsOutput struct {
	Keys                         []string
	Handled_Consolidated_Results []ConsolidatedStandingResult
}

type SimpleCarClass struct {
	ID         int    `json:"car_class_id"`
	Short_Name string `json:"short_name"`
}

type CarClassStandingListOfListsData struct {
	Base_Path       string
	Items           []SimpleCarClass
	Description     string
	Stylesheet_Path string
}

type Year struct {
	ID         int
	Short_Name string
}

type YearStandingListOfListsData struct {
	Base_Path       string
	Items           []Year
	Description     string
	Stylesheet_Path string
}

type SimpleTrack struct {
	ID         string
	Short_Name string
}

type TrackSubsessionListOfListsData struct {
	Base_Path       string
	Items           []SimpleTrack
	Description     string
	Stylesheet_Path string
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

func consolidate_subsession_result(subsession Subsession) ConsolidatedSubsessionResult {
	track := ""
	if subsession.Track.Config_Name != "N/A" {
		track = fmt.Sprintf("%s %s", subsession.Track.Track_Name, subsession.Track.Config_Name)
	} else {
		track = subsession.Track.Track_Name
	}
	license := ""
	if len(subsession.Allow_Licenses) > 0 {
		// @todo Determine if this causes issues
		license = subsession.Allow_Licenses[1].Group_Name
	}
	if len(subsession.Session_Results) > 0 {
		race_result := subsession.Session_Results[0].Results[0]
		return ConsolidatedSubsessionResult{
			subsession.Subsession_ID,
			track,
			subsession.Corners_Per_Lap,
			license,
			subsession.License_Category,
			race_result.Finish_Position,
			race_result.Finish_Position_In_Class,
			race_result.Laps_Lead,
			race_result.Laps_Complete,
			race_result.Average_Lap,
			race_result.Best_Lap_Time,
			race_result.Reason_Out,
			race_result.Champ_Points,
			race_result.Starting_Position,
			race_result.Starting_Position_In_Class,
			race_result.Car_Class_Short_Name,
			race_result.Division_Name,
			race_result.New_License_Level,
			race_result.New_CPI,
			race_result.Newi_Rating,
			race_result.Incidents,
			race_result.Car_Name,
			race_result.Aggregate_Champ_Points,
		}
	}
	var race_result = new(ConsolidatedSubsessionResult)
	race_result.Track = track
	return *race_result
}

func consolidate_head_to_head_subsession_result(subsession Subsession) ConsolidatedHeadToHeadSubsessionResult {
	track := ""
	if subsession.Track.Config_Name != "N/A" {
		track = fmt.Sprintf("%s %s", subsession.Track.Track_Name, subsession.Track.Config_Name)
	} else {
		track = subsession.Track.Track_Name
	}
	license := ""
	if len(subsession.Allow_Licenses) > 0 {
		// @todo Determine if this causes issues
		license = subsession.Allow_Licenses[1].Group_Name
	}
	if len(subsession.Session_Results) > 0 {
		return ConsolidatedHeadToHeadSubsessionResult{
			subsession.Subsession_ID,
			track,
			subsession.Corners_Per_Lap,
			license,
			subsession.License_Category,
			subsession.Event_Average_Lap,
			subsession.Event_Laps_Complete,
			subsession.Event_Strength_Of_Field,
			subsession.Num_Caution_Laps,
			subsession.Num_Cautions,
			subsession.Num_Lead_Changes,
			subsession.Race_Week_Num,
			subsession.Season_Name,
		}
	}
	var race_result = new(ConsolidatedHeadToHeadSubsessionResult)
	race_result.Track = track
	return *race_result
}

func create_seasons_function_map() template.FuncMap {
	return template.FuncMap{
		"retrieve_keys_to_display": func() []string {
			return []string{
				"subsession_id",
				"track",
				"corners_per_lap",
				"license",
				"license_category",
				"finish_position",
				"finish_position_in_class",
				"laps_lead",
				"laps_complete",
				"average_lap",
				"best_lap_time",
				"reason_out",
				"champ_points",
				"starting_position",
				"starting_position_in_class",
				"car_class_short_name",
				"division_name",
				"new_license_level",
				"new_cpi",
				"newi_rating",
				"incidents",
				"car_name",
				"aggregate_champ_points",
			}
		},
		"retrieve_key_map": func() map[string]string {
			key_map := make(map[string]string)
			key_map["subsession_id"] = "Subsession ID"
			key_map["track"] = "Track"
			key_map["corners_per_lap"] = "Corners"
			key_map["license"] = "License"
			key_map["license_category"] = "License Category"
			key_map["finish_position"] = "Finish"
			key_map["finish_position_in_class"] = "Finish (In Class)"
			key_map["laps_lead"] = "Laps Lead"
			key_map["laps_complete"] = "Laps Complete"
			key_map["average_lap"] = "AVG Laptime"
			key_map["best_lap_time"] = "Fastest Lap Time"
			key_map["reason_out"] = "Status"
			key_map["champ_points"] = "Points"
			key_map["starting_position"] = "Starting Position"
			key_map["starting_position_in_class"] = "Starting Position (In Class)"
			key_map["car_class_short_name"] = "Car Class"
			key_map["division_name"] = "Division"
			key_map["new_license_level"] = "New License Level"
			key_map["new_cpi"] = "New Corners Per Incident"
			key_map["newi_rating"] = "New iRating"
			key_map["incidents"] = "Incidents"
			key_map["car_name"] = "Car"
			key_map["aggregate_champ_points"] = "Best Points For Week"
			return key_map
		},
		"handle_results": func(keys_to_display []string, results []Subsession) HandleConsolidatedResultsOutput {
			unique_key_map := make(map[string]string)
			handled_results := []ConsolidatedSubsessionResult{}
			consolidated_subsession_results := []ConsolidatedSubsessionResult{}
			for _, subsession := range results {
				consolidated_subsession_results = append(consolidated_subsession_results, consolidate_subsession_result(subsession))
			}
			for _, result := range consolidated_subsession_results {
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
			return HandleConsolidatedResultsOutput{keys, handled_results}
		},
		"return_value_at_key": func(result ConsolidatedSubsessionResult, key string) any {
			reflected_result := reflect.ValueOf(&result).Elem()
			value := reflected_result.FieldByName(key)
			if slices.Contains([]string{"Interval", "Class_Interval", "Average_Lap", "Best_Lap_Time"}, key) {
				return convert_lap_time(int(value.Int()))
			}
			return value
		},
		"return_key_label": func(key_map map[string]string, key string) string {
			return key_map[strings.ToLower(key)]
		},
		"lowercase": func(value string) string {
			return strings.ToLower(value)
		},
		"generate_summary_stats": func(standing_with_subsessions StandingWithSubsessions) SummaryStats {
			weeks_counted := standing_with_subsessions.Standing.Season_Driver_Data.Weeks_Counted
			subsessions := standing_with_subsessions.Subsessions
			total_start_position := 0
			total_finish_position := 0
			total_incidents := 0
			wins := 0
			subsessions_with_results := slices.Collect(func(yield func(Subsession) bool) {
				for _, subsession := range subsessions {
					race_sessions := slices.Collect(func(yield func(Session) bool) {
						for _, session := range subsession.Session_Results {
							if session.Simsession_Name == "RACE" || session.Simsession_Name == "FEATURE" || session.Simsession_Name == "N/A" {
								if !yield(session) {
									return
								}
							}
						}
					})
					if len(race_sessions) == 1 {
						subsession.Session_Results = race_sessions
						race_results := slices.Collect(func(yield func(SessionResult) bool) {
							for _, race_result := range race_sessions[0].Results {
								if race_result.Cust_ID == standing_with_subsessions.Standing.Season_Driver_Data.Cust_ID && race_result.Car_Class_ID == standing_with_subsessions.Standing.Car_Class_ID {
									if !yield(race_result) {
										return
									}
								}
							}
						})
						if len(race_results) == 1 {
							subsession.Session_Results[0].Results = race_results
							if !yield(subsession) {
								return
							}
						}
					}
				}
			})
			sort.Slice(subsessions_with_results, func(i, j int) bool {
				return subsessions_with_results[i].Session_Results[0].Results[0].Champ_Points > subsessions_with_results[j].Session_Results[0].Results[0].Champ_Points
			})
			var subsessions_slice []Subsession
			if len(subsessions_with_results) > 8 {
				subsessions_slice = subsessions_with_results[0:9]
			} else {
				subsessions_slice = subsessions_with_results
			}
			for _, subsession := range subsessions_slice {
				if len(subsession.Session_Results) > 0 {
					result := subsession.Session_Results[0].Results[0]
					total_start_position += result.Starting_Position_In_Class + 1
					total_finish_position += result.Finish_Position_In_Class + 1
					if result.Finish_Position_In_Class+1 == 1 {
						wins += 1
					}
					total_incidents += result.Incidents
				}
			}
			avg_start := strconv.FormatFloat(float64(float64(total_start_position)/float64(weeks_counted)), 'f', 2, 64)
			avg_finish := strconv.FormatFloat(float64(float64(total_finish_position)/float64(weeks_counted)), 'f', 2, 64)
			return SummaryStats{
				Avg_Start:  avg_start,
				Avg_Finish: avg_finish,
				Wins:       wins,
				Incidents:  total_incidents,
			}
		},
	}
}

func consolidate_standing_result(standing Standing) ConsolidatedStandingResult {
	matcher, _ := regexp.Compile("[0-9]{4}")
	matchs := matcher.FindAllString(standing.Season_Name, 2)
	year := matchs[len(matchs)-1]
	year_as_int, _ := strconv.Atoi(year)
	matcher_2, _ := regexp.Compile("Season [0-9]{1}")
	match := matcher_2.FindString(standing.Season_Name)
	season_number := "0"
	if len(match) > 0 {
		matcher_3, _ := regexp.Compile("[0-9]{1}")
		match_2 := matcher_3.FindString(match)
		if len(match_2) > 0 {
			season_number = match_2
		}
	}
	season_number_as_int, _ := strconv.Atoi(season_number)
	return ConsolidatedStandingResult{
		standing.Season_ID,
		year_as_int,
		season_number_as_int,
		standing.Car_Class_ID,
		standing.Season_ID,
		standing.Division,
		standing.Division_Rank,
		standing.Overall_Rank,
		standing.Season_Driver_Data.Points,
		standing.Season_Name,
		standing.Season_Driver_Data.Weeks_Counted,
		standing.Season_Driver_Data.Starts,
		standing.Season_Driver_Data.Wins,
		standing.Season_Driver_Data.Top5,
		standing.Season_Driver_Data.Top25_Percent,
		standing.Season_Driver_Data.Poles,
		standing.Season_Driver_Data.Avg_Start_Position,
		standing.Season_Driver_Data.Avg_Finish_Position,
		standing.Season_Driver_Data.Avg_Field_Size,
		standing.Season_Driver_Data.Laps,
		standing.Season_Driver_Data.Laps_Led,
		standing.Season_Driver_Data.Incidents,
		standing.Season_Driver_Data.Cust_ID,
	}
}

func generate_subsession_pages() {
	raw_subsessions_input, err := os.ReadFile("./1-subsessions-output.json")
	if err != nil {
		fmt.Println(1, err)
	}
	var subsessions []Subsession
	err = json.Unmarshal(raw_subsessions_input, &subsessions)
	if err != nil {
		fmt.Println(2, err)
	}
	subsessions_function_map := template.FuncMap{
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
	// @todo Account for team sessions
	subsessions_html_template, err := template.New("subsessions.html").Funcs(subsessions_function_map).ParseFiles("subsessions.html")
	for subsession_index, subsession := range subsessions {
		fmt.Println(fmt.Sprintf("Creating %s file of %s", strconv.Itoa(subsession_index+1), strconv.Itoa(len(subsessions))))
		os.MkdirAll(fmt.Sprintf("./subsession/%s", strconv.Itoa(subsession.Subsession_ID)), os.ModePerm)
		file, err := os.Create(fmt.Sprintf("./subsession/%s/index.html", strconv.Itoa(subsession.Subsession_ID)))
		if err != nil {
			fmt.Println(3, err)
		}
		err = subsessions_html_template.Execute(file, subsession)
		if err != nil {
			fmt.Println(4, err)
		}
		file.Close()
	}
}

func generate_standing_pages() {
	raw_subsessions_input, err := os.ReadFile("./1-subsessions-output.json")
	if err != nil {
		fmt.Println(1, err)
	}
	var subsessions []Subsession
	err = json.Unmarshal(raw_subsessions_input, &subsessions)
	if err != nil {
		fmt.Println(2, err)
	}
	raw_standings_input, err := os.ReadFile("./standings-output.json")
	if err != nil {
		fmt.Println(5, err)
	}
	var standings []Standing
	err = json.Unmarshal(raw_standings_input, &standings)
	if err != nil {
		fmt.Println(6, err)
	}
	seasons_function_map := create_seasons_function_map()
	raw_past_seasons_input, err := os.ReadFile("./past-seasons-output.json")
	if err != nil {
		fmt.Println(7, raw_past_seasons_input)
	}
	var past_seasons []PastSeason
	err = json.Unmarshal(raw_past_seasons_input, &past_seasons)
	if err != nil {
		fmt.Println(8, err)
	}
	season_id_to_race_week_map := make(map[int][]RaceWeek)
	for _, past_season := range past_seasons {
		for _, season := range past_season.Series.Seasons {
			season_id_to_race_week_map[season.Season_ID] = season.Race_Weeks
		}
	}
	// @todo Account for team sessions
	season_html_template, err := template.New("seasons.html").Funcs(seasons_function_map).ParseFiles("seasons.html")
	for standings_index, standing := range standings {
		fmt.Println(fmt.Sprintf("Creating %s file of %s", strconv.Itoa(standings_index+1), strconv.Itoa(len(standings))))
		subsessions_for_season := slices.Collect(func(yield func(Subsession) bool) {
			for _, subsession := range subsessions {
				if subsession.Season_ID == standing.Season_ID {
					race_sessions := slices.Collect(func(yield func(Session) bool) {
						for _, session := range subsession.Session_Results {
							if session.Simsession_Name == "RACE" || session.Simsession_Name == "FEATURE" || session.Simsession_Name == "N/A" {
								if !yield(session) {
									return
								}
							}
						}
					})
					if len(race_sessions) == 1 {
						subsession.Session_Results = race_sessions
						race_results := slices.Collect(func(yield func(SessionResult) bool) {
							for _, race_result := range race_sessions[0].Results {
								if race_result.Cust_ID == standing.Season_Driver_Data.Cust_ID && race_result.Car_Class_ID == standing.Car_Class_ID {
									if !yield(race_result) {
										return
									}
								}
							}
						})
						if len(race_results) == 1 {
							subsession.Session_Results[0].Results = race_results
							if !yield(subsession) {
								return
							}
						}
					}
				}
			}
		})
		if len(subsessions_for_season) >= 1 {
			standing_with_subsessions := StandingWithSubsessions{standing, subsessions_for_season}
			var subsessions_by_week []Subsession
			race_weeks_for_season := season_id_to_race_week_map[standing.Season_ID]
			for _, race_week := range race_weeks_for_season {
				subsessions_for_week := slices.Collect(func(yield func(Subsession) bool) {
					for _, subsession := range subsessions_for_season {
						if subsession.Race_Week_Num == race_week.Race_Week_Num {
							if !yield(subsession) {
								return
							}
						}
					}
				})
				sort.Slice(subsessions_for_week, func(i, j int) bool {
					return subsessions_for_week[i].Session_Results[0].Results[0].Champ_Points > subsessions_for_week[j].Session_Results[0].Results[0].Champ_Points
				})
				if len(subsessions_for_week) > 0 {
					best_user_result := subsessions_for_week[0]
					subsessions_by_week = append(subsessions_by_week, best_user_result)
				} else {
					var subsession_for_week = new(Subsession)
					subsession_for_week.Track = race_week.Track
					subsessions_by_week = append(subsessions_by_week, *subsession_for_week)
				}
			}
			os.MkdirAll(fmt.Sprintf("./user/%s/season/%s/car-class/%s", strconv.Itoa(standing.Season_Driver_Data.Cust_ID), strconv.Itoa(standing.Season_ID), strconv.Itoa(standing.Car_Class_ID)), os.ModePerm)
			file, err := os.Create(fmt.Sprintf("./user/%s/season/%s/car-class/%s/index.html", strconv.Itoa(standing.Season_Driver_Data.Cust_ID), strconv.Itoa(standing.Season_ID), strconv.Itoa(standing.Car_Class_ID)))
			if err != nil {
				fmt.Println(7, err)
			}
			err = season_html_template.Execute(file, standing_with_subsessions)
			if err != nil {
				fmt.Println(8, err)
			}
			file.Close()
			season_summaries_html_template, err := template.New("season-summaries.html").Funcs(seasons_function_map).ParseFiles("season-summaries.html")
			standings_with_subsession_by_week := StandingWithSubsessions{standing, subsessions_by_week}
			os.MkdirAll(fmt.Sprintf("./user/%s/season/%s/car-class/%s/season-summary", strconv.Itoa(standing.Season_Driver_Data.Cust_ID), strconv.Itoa(standing.Season_ID), strconv.Itoa(standing.Car_Class_ID)), os.ModePerm)
			season_summary_file, err := os.Create(fmt.Sprintf("./user/%s/season/%s/car-class/%s/season-summary/index.html", strconv.Itoa(standing.Season_Driver_Data.Cust_ID), strconv.Itoa(standing.Season_ID), strconv.Itoa(standing.Car_Class_ID)))
			if err != nil {
				fmt.Println(9, err)
			}
			err = season_summaries_html_template.Execute(season_summary_file, standings_with_subsession_by_week)
			if err != nil {
				fmt.Println(10, err)
			}
			season_summary_file.Close()
		}
	}
}

func generate_subsession_list_pages() {
	raw_subsessions_input, err := os.ReadFile("./1-subsessions-output.json")
	if err != nil {
		fmt.Println(1, err)
	}
	var subsessions []Subsession
	err = json.Unmarshal(raw_subsessions_input, &subsessions)
	if err != nil {
		fmt.Println(2, err)
	}
	subsession_list_of_lists_function_map := template.FuncMap{
		"generate_full_path": func(base_path string, id int) string {
			return fmt.Sprintf("%s%s", base_path, strconv.Itoa(id))
		},
	}
	subsession_by_track_list_of_lists_function_map := template.FuncMap{
		"generate_full_path": func(base_path string, id string) string {
			return fmt.Sprintf("%s%s", base_path, id)
		},
	}
	raw_car_class_input, err := os.ReadFile("./car-classes.json")
	if err != nil {
		fmt.Println(1, err)
	}
	var car_classes []SimpleCarClass
	err = json.Unmarshal(raw_car_class_input, &car_classes)
	if err != nil {
		fmt.Println(2, err)
	}
	car_class_id_to_car_name := make(map[int]string)
	for _, car_class := range car_classes {
		car_class_id_to_car_name[car_class.ID] = car_class.Short_Name
	}
	// @todo Pull this from a file
	cust_ids := []int{300752}
	// cust_ids := []int{182407, 251134, 300752, 331322, 589449, 714312, 746377, 815162, 908575}
	for _, cust_id := range cust_ids {
		fmt.Println(fmt.Sprintf("Creating subsession listing files for %s", strconv.Itoa(cust_id)))
		subsessions_for_user_by_car_class := make(map[int][]Subsession)
		subsessions_for_user_by_track := make(map[string][]Subsession)
		subsessions_for_user_by_year := make(map[int][]Subsession)
		subsessions_for_user := slices.Collect(func(yield func(Subsession) bool) {
			for _, subsession := range subsessions {
				race_sessions := slices.Collect(func(yield func(Session) bool) {
					for _, session := range subsession.Session_Results {
						if session.Simsession_Name == "RACE" || session.Simsession_Name == "FEATURE" || session.Simsession_Name == "N/A" {
							if !yield(session) {
								return
							}
						}
					}
				})
				if len(race_sessions) == 1 {
					subsession.Session_Results = race_sessions
					race_results := slices.Collect(func(yield func(SessionResult) bool) {
						for _, race_result := range race_sessions[0].Results {
							if race_result.Cust_ID == cust_id {
								if !yield(race_result) {
									return
								}
							}
						}
					})
					if len(race_results) == 1 {
						subsession.Session_Results[0].Results = race_results
						if len(subsessions_for_user_by_year[subsession.Season_Year]) > 0 {
							subsessions_for_user_by_year[subsession.Season_Year] = append(subsessions_for_user_by_year[subsession.Season_Year], subsession)
						} else {
							subsessions_for_user_by_year[subsession.Season_Year] = []Subsession{subsession}
						}
						if len(subsessions_for_user_by_car_class[race_results[0].Car_Class_ID]) > 0 {
							subsessions_for_user_by_car_class[race_results[0].Car_Class_ID] = append(subsessions_for_user_by_car_class[race_results[0].Car_Class_ID], subsession)
						} else {
							subsessions_for_user_by_car_class[race_results[0].Car_Class_ID] = []Subsession{subsession}
						}
						// @todo See if there is more effecient way to do this
						track_slug := strings.ToLower(fmt.Sprintf("%s %s", subsession.Track.Track_Name, subsession.Track.Config_Name))
						track_slug = strings.ReplaceAll(track_slug, "[", "")
						track_slug = strings.ReplaceAll(track_slug, "]", "")
						track_slug = strings.ReplaceAll(track_slug, "(", "")
						track_slug = strings.ReplaceAll(track_slug, ")", "")
						track_slug = strings.ReplaceAll(track_slug, " ", "-")
						track_slug = strings.ReplaceAll(track_slug, "---", "-")
						track_slug = strings.ReplaceAll(track_slug, "-n/a", "")
						if len(subsessions_for_user_by_track[track_slug]) > 0 {
							subsessions_for_user_by_track[track_slug] = append(subsessions_for_user_by_track[track_slug], subsession)
						} else {
							subsessions_for_user_by_track[track_slug] = []Subsession{subsession}
						}
						if !yield(subsession) {
							return
						}
					}
				}
			}
		})
		subsessions_function_map := create_seasons_function_map()
		subsessions_list_html_template, err := template.New("subsession-lists.html").Funcs(subsessions_function_map).ParseFiles("subsession-lists.html")
		if err != nil {
			fmt.Println(11, err)
		}
		subsession_list_of_lists_html_template, err := template.New("list-of-lists.html").Funcs(subsession_list_of_lists_function_map).ParseFiles("list-of-lists.html")
		if err != nil {
			fmt.Println(12, err)
		}
		subsession_by_track_list_of_lists_html_template, err := template.New("list-of-lists.html").Funcs(subsession_by_track_list_of_lists_function_map).ParseFiles("list-of-lists.html")
		if err != nil {
			fmt.Println(12, err)
		}
		os.MkdirAll(fmt.Sprintf("./user/%s/subsessions", strconv.Itoa(cust_id)), os.ModePerm)
		file, err := os.Create(fmt.Sprintf("./user/%s/subsessions/index.html", strconv.Itoa(cust_id)))
		if err != nil {
			fmt.Println(12, err)
		}
		description := fmt.Sprintf("Subsessions - %s", strconv.Itoa(cust_id))
		sort.Slice(subsessions_for_user, func(i, j int) bool {
			return subsessions_for_user[i].Subsession_ID > subsessions_for_user[j].Subsession_ID
		})
		description_with_all_subsessions := SubsessionListSheetData{description, subsessions_for_user, "../../season.css", "../../alpine-components/table.js"}
		err = subsessions_list_html_template.Execute(file, description_with_all_subsessions)
		if err != nil {
			fmt.Println(13, err)
		}
		stylesheet_file_path := "../../../../season.css"
		table_component_file_path := "../../../../alpine-components/table.js"
		car_classes_index := make([]SimpleCarClass, len(subsessions_for_user_by_car_class))
		car_class_iteration := 0
		for key, value := range subsessions_for_user_by_car_class {
			car_classes_index[car_class_iteration] = SimpleCarClass{ID: key, Short_Name: car_class_id_to_car_name[key]}
			sort.Slice(value, func(i, j int) bool {
				return value[i].Subsession_ID > value[j].Subsession_ID
			})
			os.MkdirAll(fmt.Sprintf("./user/%s/subsessions/by-car-class/%s", strconv.Itoa(cust_id), strconv.Itoa(key)), os.ModePerm)
			file, err := os.Create(fmt.Sprintf("./user/%s/subsessions/by-car-class/%s/index.html", strconv.Itoa(cust_id), strconv.Itoa(key)))
			if err != nil {
				fmt.Println(14, err)
			}
			description := fmt.Sprintf("Subsessions list for user ID - %s and car class ID - %s", strconv.Itoa(cust_id), strconv.Itoa(key))
			description_with_subsessions := SubsessionListSheetData{description, value, stylesheet_file_path, table_component_file_path}
			err = subsessions_list_html_template.Execute(file, description_with_subsessions)
			if err != nil {
				fmt.Println(15, err)
			}
			car_class_iteration++
		}
		sort.Slice(car_classes_index, func(i, j int) bool {
			return car_classes_index[i].Short_Name < car_classes_index[j].Short_Name
		})
		car_class_index_file, err := os.Create(fmt.Sprintf("./user/%s/subsessions/by-car-class/index.html", strconv.Itoa(cust_id)))
		if err != nil {
			fmt.Println(16, err)
		}
		car_class_index_file_description := fmt.Sprintf("Subsessions by car class page for user ID: %s", strconv.Itoa(cust_id))
		car_class_base_path := fmt.Sprintf("/Stat-N-Track/user/%s/subsessions/by-car-class/", strconv.Itoa(cust_id))
		car_class_standings_list_of_lists_data := CarClassStandingListOfListsData{car_class_base_path, car_classes_index, car_class_index_file_description, "../../../season.css"}
		err = subsession_list_of_lists_html_template.Execute(car_class_index_file, car_class_standings_list_of_lists_data)
		if err != nil {
			fmt.Println(16, err)
		}
		tracks_index := make([]SimpleTrack, len(subsessions_for_user_by_track))
		tracks_iteration := 0
		for key, value := range subsessions_for_user_by_track {
			config_name := ""
			if value[0].Track.Config_Name != "N/A" {
				config_name = value[0].Track.Config_Name
			}
			tracks_index[tracks_iteration] = SimpleTrack{ID: key, Short_Name: fmt.Sprintf("%s %s", value[0].Track.Track_Name, config_name)}
			sort.Slice(value, func(i, j int) bool {
				return value[i].Subsession_ID > value[j].Subsession_ID
			})
			os.MkdirAll(fmt.Sprintf("./user/%s/subsessions/by-track/%s", strconv.Itoa(cust_id), key), os.ModePerm)
			file, err := os.Create(fmt.Sprintf("./user/%s/subsessions/by-track/%s/index.html", strconv.Itoa(cust_id), key))
			if err != nil {
				fmt.Println(16, err)
			}
			description := fmt.Sprintf("Subsessions list for user ID - %s and track - %s", strconv.Itoa(cust_id), key)
			description_with_subsessions := SubsessionListSheetData{description, value, stylesheet_file_path, table_component_file_path}
			err = subsessions_list_html_template.Execute(file, description_with_subsessions)
			if err != nil {
				fmt.Println(17, err)
			}
			tracks_iteration++
		}
		sort.Slice(tracks_index, func(i, j int) bool {
			return tracks_index[i].Short_Name < tracks_index[j].Short_Name
		})
		track_index_file, err := os.Create(fmt.Sprintf("./user/%s/subsessions/by-track/index.html", strconv.Itoa(cust_id)))
		if err != nil {
			fmt.Println(16, err)
		}
		track_index_file_description := fmt.Sprintf("Subsessions by track page for user ID: %s", strconv.Itoa(cust_id))
		track_base_path := fmt.Sprintf("/Stat-N-Track/user/%s/subsessions/by-track/", strconv.Itoa(cust_id))
		track_subsession_list_of_lists_data := TrackSubsessionListOfListsData{track_base_path, tracks_index, track_index_file_description, "../../../season.css"}
		err = subsession_by_track_list_of_lists_html_template.Execute(track_index_file, track_subsession_list_of_lists_data)
		if err != nil {
			fmt.Println(16, err)
		}
		years_index := make([]Year, len(subsessions_for_user_by_year))
		years_iteration := 0
		for key, value := range subsessions_for_user_by_year {
			years_index[years_iteration] = Year{ID: key, Short_Name: strconv.Itoa(key)}
			sort.Slice(value, func(i, j int) bool {
				return value[i].Subsession_ID > value[j].Subsession_ID
			})
			os.MkdirAll(fmt.Sprintf("./user/%s/subsessions/by-year/%s", strconv.Itoa(cust_id), strconv.Itoa(key)), os.ModePerm)
			file, err := os.Create(fmt.Sprintf("./user/%s/subsessions/by-year/%s/index.html", strconv.Itoa(cust_id), strconv.Itoa(key)))
			if err != nil {
				fmt.Println(18, err)
			}
			description := fmt.Sprintf("Subsessions list for user ID - %s and year - %s", strconv.Itoa(cust_id), strconv.Itoa(key))
			description_with_subsessions := SubsessionListSheetData{description, value, stylesheet_file_path, table_component_file_path}
			err = subsessions_list_html_template.Execute(file, description_with_subsessions)
			if err != nil {
				fmt.Println(19, err)
			}
			years_iteration++
		}
		sort.Slice(years_index, func(i, j int) bool {
			return years_index[i].ID < years_index[j].ID
		})
		year_index_file, err := os.Create(fmt.Sprintf("./user/%s/subsessions/by-year/index.html", strconv.Itoa(cust_id)))
		if err != nil {
			fmt.Println(16, err)
		}
		year_index_file_description := fmt.Sprintf("Subsessions by year page for user ID: %s", strconv.Itoa(cust_id))
		year_base_path := fmt.Sprintf("/Stat-N-Track/user/%s/subsessions/by-year/", strconv.Itoa(cust_id))
		year_standings_list_of_lists_data := YearStandingListOfListsData{year_base_path, years_index, year_index_file_description, "../../../season.css"}
		err = subsession_list_of_lists_html_template.Execute(year_index_file, year_standings_list_of_lists_data)
		if err != nil {
			fmt.Println(16, err)
		}
	}
}

func generate_standing_list_pages() {
	raw_standings_input, err := os.ReadFile("./standings-output.json")
	if err != nil {
		fmt.Println(5, err)
	}
	var standings []Standing
	err = json.Unmarshal(raw_standings_input, &standings)
	if err != nil {
		fmt.Println(6, err)
	}
	standings_list_function_map := template.FuncMap{
		"retrieve_keys_to_display": func() []string {
			return []string{
				"season_id",
				"year",
				"season_number",
				"car_class_id",
				"season_summary",
				"division",
				"division_rank",
				"overall_rank",
				"points",
				"season_name",
				"weeks_counted",
				"starts",
				"wins",
				"top5",
				"top25_percent",
				"poles",
				"avg_start_position",
				"avg_finish_position",
				"avg_field_size",
				"laps",
				"laps_led",
				"incidents",
			}
		},
		"retrieve_key_map": func() map[string]string {
			key_map := make(map[string]string)
			key_map["season_id"] = "Season ID"
			key_map["year"] = "Year"
			key_map["season_number"] = "Season Number"
			key_map["car_class_id"] = "Car Class ID"
			key_map["season_summary"] = "Season Summary"
			key_map["division"] = "Division"
			key_map["division_rank"] = "Rank (Division)"
			key_map["overall_rank"] = "Rank (Overall)"
			key_map["points"] = "Points"
			key_map["season_name"] = "Season Name"
			key_map["weeks_counted"] = "Participation Weeks"
			key_map["starts"] = "Starts"
			key_map["wins"] = "Wins"
			key_map["top5"] = "Top 5s"
			key_map["top25_percent"] = "Top 25% Finishes"
			key_map["poles"] = "Poles"
			key_map["avg_start_position"] = "Average Starting Position"
			key_map["avg_finish_position"] = "Average Finishing Position"
			key_map["avg_field_size"] = "Average Field Size"
			key_map["laps"] = "Laps"
			key_map["laps_led"] = "Laps Led"
			key_map["incidents"] = "Incidents"
			return key_map
		},
		"handle_results": func(keys_to_display []string, results []Standing) HandleConsolidatedStandingResultsOutput {
			unique_key_map := make(map[string]string)
			handled_results := []ConsolidatedStandingResult{}
			consolidated_standing_results := []ConsolidatedStandingResult{}
			for _, standing := range results {
				consolidated_standing_results = append(consolidated_standing_results, consolidate_standing_result(standing))
			}
			for _, result := range consolidated_standing_results {
				reflected_result := reflect.ValueOf(&result).Elem()
				fields := reflect.VisibleFields(reflect.TypeOf(result))
				for _, field := range fields {
					if slices.Contains(keys_to_display, strings.ToLower(field.Name)) {
						unique_key_map[strings.ToLower(field.Name)] = field.Name
						// Being lazy because types are cumbersome sometimes
						if field.Name == "Division" {
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
			return HandleConsolidatedStandingResultsOutput{keys, handled_results}
		},
		"return_value_at_key": func(result ConsolidatedStandingResult, key string) any {
			reflected_result := reflect.ValueOf(&result).Elem()
			return reflected_result.FieldByName(key)
		},
		"return_key_label": func(key_map map[string]string, key string) string {
			return key_map[strings.ToLower(key)]
		},
		"lowercase": func(value string) string {
			return strings.ToLower(value)
		},
		"should_have_link": func(key string) bool {
			lowercase_key := strings.ToLower(key)
			if lowercase_key == "season_id" || lowercase_key == "year" || lowercase_key == "car_class_id" || lowercase_key == "season_summary" {
				return true
			}
			return false
		},
		"generate_link_for_key": func(key string, consolidated_standing_result ConsolidatedStandingResult) string {
			lowercase_key := strings.ToLower(key)
			cust_id_as_string := strconv.Itoa(consolidated_standing_result.Cust_ID)
			season_id_as_string := strconv.Itoa(consolidated_standing_result.Season_ID)
			year_as_string := strconv.Itoa(consolidated_standing_result.Year)
			car_class_id_as_string := strconv.Itoa(consolidated_standing_result.Car_Class_ID)
			if lowercase_key == "season_id" {
				return fmt.Sprintf("/Stat-N-Track/user/%s/season/%s/car-class/%s/", cust_id_as_string, season_id_as_string, car_class_id_as_string)
			}
			if lowercase_key == "year" {
				return fmt.Sprintf("/Stat-N-Track/user/%s/standings/by-year/%s/", cust_id_as_string, year_as_string)
			}
			if lowercase_key == "car_class_id" {
				return fmt.Sprintf("/Stat-N-Track/user/%s/standings/by-car-class/%s/", cust_id_as_string, car_class_id_as_string)
			}
			if lowercase_key == "season_summary" {
				return fmt.Sprintf("/Stat-N-Track/user/%s/season/%s/car-class/%s/season-summary/", cust_id_as_string, season_id_as_string, car_class_id_as_string)
			}
			return ""
		},
	}
	standings_list_of_lists_function_map := template.FuncMap{
		"generate_full_path": func(base_path string, id int) string {
			return fmt.Sprintf("%s%s", base_path, strconv.Itoa(id))
		},
	}
	raw_car_class_input, err := os.ReadFile("./car-classes.json")
	if err != nil {
		fmt.Println(1, err)
	}
	var car_classes []SimpleCarClass
	err = json.Unmarshal(raw_car_class_input, &car_classes)
	if err != nil {
		fmt.Println(2, err)
	}
	car_class_id_to_car_name := make(map[int]string)
	for _, car_class := range car_classes {
		car_class_id_to_car_name[car_class.ID] = car_class.Short_Name
	}
	// @todo Pull this from a file
	cust_ids := []int{300752}
	// cust_ids := []int{182407, 251134, 300752, 331322, 589449, 714312, 746377, 815162, 908575}
	for _, cust_id := range cust_ids {
		fmt.Println(fmt.Sprintf("Creating standing listing files for %s", strconv.Itoa(cust_id)))
		standings_for_user_by_car_class := make(map[int][]Standing)
		standings_for_user_by_year := make(map[int][]Standing)
		standings_for_user_full_participation := []Standing{}
		standings_for_user := slices.Collect(func(yield func(Standing) bool) {
			for _, standing := range standings {
				if standing.Season_Driver_Data.Cust_ID == cust_id {
					if len(standings_for_user_by_car_class[standing.Car_Class_ID]) > 0 {
						standings_for_user_by_car_class[standing.Car_Class_ID] = append(standings_for_user_by_car_class[standing.Car_Class_ID], standing)
					} else {
						standings_for_user_by_car_class[standing.Car_Class_ID] = []Standing{standing}
					}
					matcher, _ := regexp.Compile("[0-9]{4}")
					matchs := matcher.FindAllString(standing.Season_Name, 2)
					year := matchs[len(matchs)-1]
					year_as_int, _ := strconv.Atoi(year)
					if len(standings_for_user_by_year[year_as_int]) >= 0 {
						standings_for_user_by_year[year_as_int] = append(standings_for_user_by_year[year_as_int], standing)
					} else {
						standings_for_user_by_year[year_as_int] = []Standing{standing}
					}
					if standing.Season_Driver_Data.Weeks_Counted >= 8 {
						standings_for_user_full_participation = append(standings_for_user_full_participation, standing)
					}
					if !yield(standing) {
						return
					}
				}
			}
		})
		standings_list_html_template, err := template.New("standing-lists.html").Funcs(standings_list_function_map).ParseFiles("standing-lists.html")
		if err != nil {
			fmt.Println(11, err)
		}
		standings_list_of_lists_html_template, err := template.New("list-of-lists.html").Funcs(standings_list_of_lists_function_map).ParseFiles("list-of-lists.html")
		if err != nil {
			fmt.Println(12, err)
		}
		os.MkdirAll(fmt.Sprintf("./user/%s/standings", strconv.Itoa(cust_id)), os.ModePerm)
		file, err := os.Create(fmt.Sprintf("./user/%s/standings/index.html", strconv.Itoa(cust_id)))
		if err != nil {
			fmt.Println(12, err)
		}
		description := fmt.Sprintf("All Standings - %s", strconv.Itoa(cust_id))
		sort.Slice(standings_for_user, func(i, j int) bool {
			return standings_for_user[i].Season_ID > standings_for_user[j].Season_ID
		})
		description_with_all_standings := StandingListSheetData{description, standings_for_user, "../../season.css", "../../alpine-components/table.js"}
		err = standings_list_html_template.Execute(file, description_with_all_standings)
		if err != nil {
			fmt.Println(13, err)
		}
		sort.Slice(standings_for_user_full_participation, func(i, j int) bool {
			return standings_for_user_full_participation[i].Season_ID > standings_for_user_full_participation[j].Season_ID
		})
		os.MkdirAll(fmt.Sprintf("./user/%s/standings/full-participation", strconv.Itoa(cust_id)), os.ModePerm)
		full_participation_file, err := os.Create(fmt.Sprintf("./user/%s/standings/full-participation/index.html", strconv.Itoa(cust_id)))
		if err != nil {
			fmt.Println(16, err)
		}
		full_participation_description := fmt.Sprintf("Fully participated standings list for user ID - %s", strconv.Itoa(cust_id))
		description_with_full_participation_standings := StandingListSheetData{full_participation_description, standings_for_user_full_participation, "../../../season.css", "../../../alpine-components/table.js"}
		err = standings_list_html_template.Execute(full_participation_file, description_with_full_participation_standings)
		if err != nil {
			fmt.Println(17, err)
		}
		stylesheet_file_path := "../../../../season.css"
		table_component_file_path := "../../../../alpine-components/table.js"
		car_classes_index := make([]SimpleCarClass, len(standings_for_user_by_car_class))
		car_class_iteration := 0
		for key, value := range standings_for_user_by_car_class {
			car_classes_index[car_class_iteration] = SimpleCarClass{ID: key, Short_Name: car_class_id_to_car_name[key]}
			sort.Slice(value, func(i, j int) bool {
				return value[i].Season_ID > value[j].Season_ID
			})
			os.MkdirAll(fmt.Sprintf("./user/%s/standings/by-car-class/%s", strconv.Itoa(cust_id), strconv.Itoa(key)), os.ModePerm)
			file, err := os.Create(fmt.Sprintf("./user/%s/standings/by-car-class/%s/index.html", strconv.Itoa(cust_id), strconv.Itoa(key)))
			if err != nil {
				fmt.Println(14, err)
			}
			description := fmt.Sprintf("Standings list for user ID - %s and car class ID - %s", strconv.Itoa(cust_id), strconv.Itoa(key))
			description_with_standings := StandingListSheetData{description, value, stylesheet_file_path, table_component_file_path}
			err = standings_list_html_template.Execute(file, description_with_standings)
			if err != nil {
				fmt.Println(15, err)
			}
			car_class_iteration++
		}
		sort.Slice(car_classes_index, func(i, j int) bool {
			return car_classes_index[i].Short_Name < car_classes_index[j].Short_Name
		})
		car_class_index_file, err := os.Create(fmt.Sprintf("./user/%s/standings/by-car-class/index.html", strconv.Itoa(cust_id)))
		if err != nil {
			fmt.Println(16, err)
		}
		car_class_index_file_description := fmt.Sprintf("Car class standings page for user ID: %s", strconv.Itoa(cust_id))
		car_class_base_path := fmt.Sprintf("/Stat-N-Track/user/%s/standings/by-car-class/", strconv.Itoa(cust_id))
		car_class_standings_list_of_lists_data := CarClassStandingListOfListsData{car_class_base_path, car_classes_index, car_class_index_file_description, "../../../season.css"}
		err = standings_list_of_lists_html_template.Execute(car_class_index_file, car_class_standings_list_of_lists_data)
		if err != nil {
			fmt.Println(16, err)
		}
		years_index := make([]Year, len(standings_for_user_by_year))
		years_iteration := 0
		for key, value := range standings_for_user_by_year {
			years_index[years_iteration] = Year{ID: key, Short_Name: strconv.Itoa(key)}
			sort.Slice(value, func(i, j int) bool {
				return value[i].Season_ID > value[j].Season_ID
			})
			os.MkdirAll(fmt.Sprintf("./user/%s/standings/by-year/%s", strconv.Itoa(cust_id), strconv.Itoa(key)), os.ModePerm)
			file, err := os.Create(fmt.Sprintf("./user/%s/standings/by-year/%s/index.html", strconv.Itoa(cust_id), strconv.Itoa(key)))
			if err != nil {
				fmt.Println(18, err)
			}
			description := fmt.Sprintf("Standings list for user ID - %s and year - %s", strconv.Itoa(cust_id), strconv.Itoa(key))
			description_with_standings := StandingListSheetData{description, value, stylesheet_file_path, table_component_file_path}
			err = standings_list_html_template.Execute(file, description_with_standings)
			if err != nil {
				fmt.Println(19, err)
			}
			years_iteration++
		}
		sort.Slice(years_index, func(i, j int) bool {
			return years_index[i].ID < years_index[j].ID
		})
		year_index_file, err := os.Create(fmt.Sprintf("./user/%s/standings/by-year/index.html", strconv.Itoa(cust_id)))
		if err != nil {
			fmt.Println(16, err)
		}
		year_index_file_description := fmt.Sprintf("Standings by year page for user ID: %s", strconv.Itoa(cust_id))
		year_base_path := fmt.Sprintf("/Stat-N-Track/user/%s/standings/by-year/", strconv.Itoa(cust_id))
		year_standings_list_of_lists_data := YearStandingListOfListsData{year_base_path, years_index, year_index_file_description, "../../../season.css"}
		err = standings_list_of_lists_html_template.Execute(year_index_file, year_standings_list_of_lists_data)
		if err != nil {
			fmt.Println(16, err)
		}
	}
}

func generate_head_to_head_pages() {
	raw_subsessions_input, err := os.ReadFile("./1-subsessions-output.json")
	if err != nil {
		fmt.Println(1, err)
	}
	var subsessions []Subsession
	err = json.Unmarshal(raw_subsessions_input, &subsessions)
	if err != nil {
		fmt.Println(2, err)
	}
	cust_id_pairs := []string{
		"300752_182407",
		"300752_251134",
		"300752_331322",
		"300752_589449",
		"300752_714312",
		"300752_746377",
		"300752_815162",
		"300752_908575",
		"331322_182407",
		"331322_251134",
		"331322_589449",
		"331322_714312",
		"331322_746377",
		"331322_815162",
		"331322_908575",
		"589449_182407",
		"589449_251134",
		"589449_714312",
		"589449_746377",
		"589449_815162",
		"589449_908575",
		"714312_182407",
		"714312_251134",
		"714312_746377",
		"714312_815162",
		"714312_908575",
		"746377_182407",
		"746377_251134",
		"746377_815162",
		"746377_908575",
		"815162_182407",
		"815162_251134",
		"815162_908575",
		"908575_182407",
		"908575_251134",
	}
	// fmt.Println(fmt.Sprintf("Creating %s file of %s", strconv.Itoa(standings_index+1), strconv.Itoa(len(standings))))
	subsessions_function_map := template.FuncMap{
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
		"retrieve_key_map": func() map[string]string {
			key_map := make(map[string]string)
			key_map["season_id"] = "Season ID"
			key_map["year"] = "Year"
			key_map["season_number"] = "Season Number"
			key_map["car_class_id"] = "Car Class ID"
			key_map["season_summary"] = "Season Summary"
			key_map["division"] = "Division"
			key_map["division_rank"] = "Rank (Division)"
			key_map["overall_rank"] = "Rank (Overall)"
			key_map["points"] = "Points"
			key_map["season_name"] = "Season Name"
			key_map["weeks_counted"] = "Participation Weeks"
			key_map["starts"] = "Starts"
			key_map["wins"] = "Wins"
			key_map["top5"] = "Top 5s"
			key_map["top25_percent"] = "Top 25% Finishes"
			key_map["poles"] = "Poles"
			key_map["avg_start_position"] = "Average Starting Position"
			key_map["avg_finish_position"] = "Average Finishing Position"
			key_map["avg_field_size"] = "Average Field Size"
			key_map["laps"] = "Laps"
			key_map["laps_led"] = "Laps Led"
			key_map["incidents"] = "Incidents"
			return key_map
		},
		"table_retrieve_keys_to_display": func() []string {
			return []string{
				"subsession_id",
				"track",
				"license",
				"license_category",
				"event_average_lap",
				"event_laps_complete",
				"event_strength_of_field",
				"num_caution_laps",
				"num_cautions",
				"num_lead_changes",
				"race_week_num",
				"season_name",
			}
		},
		"table_retrieve_key_map": func() map[string]string {
			key_map := make(map[string]string)
			key_map["subsession_id"] = "Subsession ID"
			key_map["track"] = "Track"
			key_map["license"] = "License"
			key_map["license_category"] = "License Category"
			key_map["event_average_lap"] = "Average Laptime"
			key_map["event_laps_complete"] = "Total Laps"
			key_map["event_strength_of_field"] = "Strength of Field"
			key_map["num_caution_laps"] = "Number of Caution Laps"
			key_map["num_cautions"] = "Number of Cautions"
			key_map["num_lead_changes"] = "Number of Lead Changes"
			key_map["race_week_num"] = "Race Week Number"
			key_map["season_name"] = "Season Name"
			return key_map
		},
		"table_handle_results": func(keys_to_display []string, results []Subsession) HandleConsolidatedHeadToHeadResultsOutput {
			unique_key_map := make(map[string]string)
			handled_results := []ConsolidatedHeadToHeadSubsessionResult{}
			consolidated_subsession_results := []ConsolidatedHeadToHeadSubsessionResult{}
			for _, subsession := range results {
				consolidated_subsession_results = append(consolidated_subsession_results, consolidate_head_to_head_subsession_result(subsession))
			}
			for _, result := range consolidated_subsession_results {
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
			return HandleConsolidatedHeadToHeadResultsOutput{keys, handled_results}
		},
		"table_return_value_at_key": func(result ConsolidatedHeadToHeadSubsessionResult, key string) any {
			reflected_result := reflect.ValueOf(&result).Elem()
			value := reflected_result.FieldByName(key)
			if slices.Contains([]string{"Interval", "Class_Interval", "Average_Lap", "Best_Lap_Time"}, key) {
				return convert_lap_time(int(value.Int()))
			}
			return value
		},
		"return_key_label": func(key_map map[string]string, key string) string {
			return key_map[strings.ToLower(key)]
		},
		"lowercase": func(value string) string {
			return strings.ToLower(value)
		},
		"generate_toggle_data": func(subsessions []Subsession) string {
			values := []string{"{\"previouslyToggled\": '',"}
			for _, subsession := range subsessions {
				values = append(values, fmt.Sprintf("\"toggle%s\": false,", strconv.Itoa(subsession.Subsession_ID)))
			}
			values = append(values, "}")
			return strings.Join(values, "")
		},
		"generate_head_to_head_wins_table": func(subsessions []Subsession) []HeadToHeadWinsOutput {
			winner_name_to_win_count_map := make(map[string]int)
			for _, subsession := range subsessions {
				winner := subsession.Session_Results[0].Results[0].Display_Name
				if winner_name_to_win_count_map[winner] >= 0 {
					winner_name_to_win_count_map[winner] += 1
				} else {
					winner_name_to_win_count_map[winner] = 1
				}
			}
			output := []HeadToHeadWinsOutput{}

			for k, v := range winner_name_to_win_count_map {
				output = append(output, HeadToHeadWinsOutput{k, v})
			}
			sort.Slice(output, func(i, j int) bool {
				return output[i].Win_Count > output[j].Win_Count
			})
			return output
		},
	}
	shared_subsessions_html_template, err := template.New("shared-subsessions.html").Funcs(subsessions_function_map).ParseFiles("shared-subsessions.html")
	for _, cust_id_pair := range cust_id_pairs {
		cust_ids := strings.Split(cust_id_pair, "_")
		head_to_head_subsessions := slices.Collect(func(yield func(Subsession) bool) {
			for _, subsession := range subsessions {
				race_sessions := slices.Collect(func(yield func(Session) bool) {
					for _, session := range subsession.Session_Results {
						if session.Simsession_Name == "RACE" || session.Simsession_Name == "FEATURE" || session.Simsession_Name == "N/A" {
							if !yield(session) {
								return
							}
						}
					}
				})
				if len(race_sessions) == 1 {
					subsession.Session_Results = race_sessions
					race_results := slices.Collect(func(yield func(SessionResult) bool) {
						for _, race_result := range race_sessions[0].Results {
							for _, cust_id := range cust_ids {
								cust_id_as_int, _ := strconv.Atoi(cust_id)
								if race_result.Cust_ID == cust_id_as_int {
									if !yield(race_result) {
										return
									}
								}
							}
						}
					})
					if len(race_results) == len(cust_ids) {
						subsession.Session_Results[0].Results = race_results
						if !yield(subsession) {
							return
						}
					}
				}
			}
		})
		os.MkdirAll(fmt.Sprintf("./user/%s/head-to-head/%s", cust_ids[0], cust_ids[1]), os.ModePerm)
		os.MkdirAll(fmt.Sprintf("./user/%s/head-to-head/%s", cust_ids[1], cust_ids[0]), os.ModePerm)
		sort.Slice(head_to_head_subsessions, func(i, j int) bool {
			return head_to_head_subsessions[i].Subsession_ID > head_to_head_subsessions[j].Subsession_ID
		})
		fmt.Println(len(head_to_head_subsessions))
		description_with_head_to_head_subsessions := SubsessionListSheetData{"HOOPLAH", head_to_head_subsessions, "../../../season.css", "../../../alpine-components/table.js"}
		file_one, err := os.Create(fmt.Sprintf("./user/%s/head-to-head/%s/index.html", cust_ids[0], cust_ids[1]))
		if err != nil {
			fmt.Println(3, err)
		}
		err = shared_subsessions_html_template.Execute(file_one, description_with_head_to_head_subsessions)
		if err != nil {
			fmt.Println(4, err)
		}
		file_one.Close()
		file_two, err := os.Create(fmt.Sprintf("./user/%s/head-to-head/%s/index.html", cust_ids[1], cust_ids[0]))
		if err != nil {
			fmt.Println(3, err)
		}
		err = shared_subsessions_html_template.Execute(file_two, description_with_head_to_head_subsessions)
		if err != nil {
			fmt.Println(4, err)
		}
	}
}

func main() {
	// generate_subsession_pages()
	// generate_standing_pages()
	// generate_subsession_list_pages()
	// generate_standing_list_pages()
	generate_head_to_head_pages()
}
