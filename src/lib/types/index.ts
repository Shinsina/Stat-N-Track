export type Car = {
  _id: number;
  ai_enabled: boolean;
  allow_number_colors: boolean;
  allow_number_font: boolean;
  allow_sponsor1: boolean;
  allow_sponsor2: boolean;
  allow_wheel_color: boolean;
  award_exempt: boolean;
  car_dirpath: string;
  car_id: number;
  car_name: string;
  car_name_abbreviated: string;
  car_types: Array<{ car_type: string }>;
  car_weight: number;
  categories: Array<string>;
  created: string;
  first_sale: string;
  forum_url: string;
  free_with_subscription: boolean;
  has_headlights: boolean;
  has_multiple_dry_tires_types: boolean;
  hp: number;
  is_ps_purchasable: boolean;
  max_power_adjust_pct: number;
  max_weight_penalty_kg: number;
  min_power_adjust_pct: number;
  package_id: number;
  patterns: number;
  price: number;
  price_display: string;
  retired: boolean;
  search_filters: string;
  sku: number;
};

export type CarClass = {
  car_class_id: number;
  cars_in_class: Array<{ car_id: number; car_name: string | undefined }>;
  name: string;
  short_name: string;
};

export type CarClassMapping = {
  car_class_id: number;
  car_class_short_name: string;
};

export type Helmet = {
  pattern: number;
  color1: string;
  color2: string;
  color3: string;
  face_type: number;
  helmet_type: number;
};

export type License = {
  parent_id: number;
  license_group: number;
  min_license_level: number;
  max_license_level: number;
  group_name: string;
};

export type Livery = {
  car_id: number;
  pattern: number;
  color1: string;
  color2: string;
  color3: string;
  number_font: number;
  number_color1: string;
  number_color2: string;
  number_color3: string;
  number_slant: number;
  sponsor1: number;
  sponsor2: number;
  car_number: string;
  wheel_color: string | null;
  rim_type: number;
};

export type RaceSummary = {
  subsession_id: number;
  average_lap: number;
  laps_complete: number;
  num_cautions: number;
  num_caution_laps: number;
  num_lead_changes: number;
  field_strength: number;
  num_opt_laps: number;
  has_opt_path: boolean;
  special_event_type: number;
  special_event_type_text: string;
};

export type RaceWeek = {
  season_id: number;
  race_week_num: number;
  series_id: number;
  series_name: string;
  season_name: string;
  schedule_name: string;
  start_date: string;
  simulated_time_multiplier: number;
  race_lap_limit: number | null;
  race_time_limit: number | null;
  start_type: string;
  restart_type: string;
  qual_attached: boolean;
  full_course_cautions: boolean;
  special_event_type: unknown;
  start_zone: boolean;
  enable_pitlane_collisions: boolean;
  short_parade_lap: boolean;
  track: Track;
  weather: Weather;
  track_state: TrackState;
  race_time_descriptors: Array<unknown>;
  car_restrictions: Array<unknown>;
  race_week_cars: Array<unknown>;
};

export type SchedulingSeason = {
  _id: number;
  active: boolean;
  allowed_season_members: unknown;
  car_class_ids: Array<number>;
  car_switching: boolean;
  car_types: Array<{ car_type: string }>;
  caution_laps_do_not_count: boolean;
  complete: boolean;
  cross_license: boolean;
  driver_change_rule: number;
  driver_changes: boolean;
  drops: number;
  enable_pitlane_collisions: boolean;
  fixed_setup: boolean;
  green_white_checkered_limit: number;
  grid_by_class: boolean;
  hardcore_level: number;
  ignore_license_for_practice: boolean;
  incident_limit: number;
  incidient_warn_mode: number;
  incident_warn_param1: number;
  incident_warn_param2: number;
  is_heat_racing: boolean;
  license_group: number;
  license_group_types: Array<{ license_group_type: number }>;
  lucky_dog: boolean;
  max_team_drivers: number;
  max_weeks: number;
  min_team_drivers: number;
  multiclass: boolean;
  must_user_diff_tire_types_in_race: boolean;
  next_race_session: unknown;
  num_opt_laps: number;
  official: boolean;
  op_duration: number;
  open_practice_session_type_id: number;
  qualifier_must_start_race: boolean;
  race_week: number;
  race_week_to_make_division: number;
  reg_user_count: number;
  region_competition: boolean;
  restrict_by_member: boolean;
  restrict_to_car: boolean;
  restrict_viewing: boolean;
  schedule_description: string;
  schedules: Array<RaceWeek>;
  season_id: number;
  season_name: string;
  season_quarter: number;
  season_short_name: string;
  season_year: number;
  send_to_open_practice: boolean;
  series_id: number;
  short_parade_lap: boolean;
  start_date: string;
  start_on_qual_tire: boolean;
  start_zone: boolean;
  track_types: Array<{ track_type: string }>;
  unsport_conduct_rule_mode: number;
};

export type SchedulingTableResult = {
  "Season Name": string;
  "License Level": string;
  carClassData: Array<CarClass>;
  tracks: Array<Track>;
  ownedTrackWeeks: Array<number>;
  "Own Car": string;
  "Overall Track Percentage Owned": number;
};

export type Season = {
  _id: string;
  car_class_id: string;
  division: number;
  overall_rank: number;
  season_driver_data: SeasonDriverData;
  season_id: string;
  season_name: string;
  division_rank: number;
};

export type SeasonDriverData = {
  rank: number;
  cust_id: number;
  display_name: number;
  division: number;
  club_id: number;
  club_name: string;
  license: SeasonDriverLicense;
  helmet: Helmet;
  weeks_counted: number;
  starts: number;
  wins: number;
  top5: number;
  top25_percent: number;
  poles: number;
  avg_start_position: number;
  avg_finish_position: number;
  avg_field_size: number;
  laps: number;
  laps_led: number;
  incidents: number;
  points: number;
  raw_points: number;
  week_dropped: boolean;
  country_code: string;
  country: string;
};

export type SeasonDriverLicense = {
  category_id: number;
  category: string;
  license_level: number;
  safety_rating: number;
  irating: number;
  color: string;
  group_name: string;
  group_id: number;
};

export type Session = {
  simsession_number: number;
  simsession_type: number;
  simsession_type_name: string;
  simsession_subtype: number;
  simsession_name: string;
  results: Array<SessionResult>;
};

export type SessionResult = {
  cust_id: number;
  display_name: string;
  finish_position: number;
  finish_position_in_class: number;
  laps_lead: number;
  laps_complete: number;
  opt_laps_complete: number;
  interval: number;
  class_interval: number;
  average_lap: number;
  best_lap_num: number;
  best_lap_time: number;
  best_nlaps_num: number;
  best_nlaps_time: number;
  best_qual_lap_at: string;
  best_qual_lap_num: number;
  best_qual_lap_time: number;
  reason_out_id: number;
  reason_out: string;
  champ_points: number;
  drop_race: boolean;
  club_points: number;
  position: number;
  qual_lap_time: number;
  starting_position: number;
  starting_position_in_class: number;
  car_class_id: number;
  car_class_name: string;
  car_class_short_name: string;
  club_id: number;
  club_name: string;
  club_shortname: string;
  division: number;
  division_name: string;
  old_license_level: number;
  old_sub_level: number;
  old_cpi: number;
  oldi_rating: number;
  old_ttrating: number;
  new_license_level: number;
  new_sub_level: number;
  new_cpi: number;
  newi_rating: number;
  new_ttrating: number;
  multiplier: number;
  license_change_oval: number;
  license_change_road: number;
  incidents: number;
  max_pct_fuel_fill: number;
  weight_penalty_kg: number;
  league_points: number;
  league_agg_points: number;
  car_id: number;
  car_name: string;
  aggregate_champ_points: number;
  livery: Livery;
  suit: Suit;
  helmet: Helmet;
  watched: boolean;
  friend: boolean;
  ai: boolean;
};

export type Subsession = {
  _id: number;
  allowed_licenses: Array<License>;
  associated_subsession_ids: Array<number>;
  can_protest: boolean;
  car_classes: Array<CarClass>;
  caution_type: number;
  cooldown_minutes: number;
  corners_per_lap: number;
  damage_model: number;
  driver_change_param1: number;
  driver_change_param2: number;
  driver_change_rule: number;
  driver_changes: boolean;
  end_time: string;
  event_average_lap: number;
  event_laps_complete: number;
  event_strength_of_field: number;
  event_type: number;
  event_type_name: string;
  heat_info_id: number;
  license_category: string;
  license_category_id: number;
  limit_minutes: number;
  max_team_drivers: number;
  max_weeks: number;
  min_team_drivers: number;
  num_caution_laps: number;
  num_cautions: number;
  num_laps_for_qual_average: number;
  num_laps_for_solo_average: number;
  num_lead_changes: number;
  official_session: boolean;
  points_type: string;
  private_session_id: number;
  race_summary: RaceSummary;
  race_week_num: number;
  results_restricted: boolean;
  season_id: number;
  season_name: string;
  season_quarter: number;
  season_short_name: string;
  season_year: number;
  series_id: number;
  series_logo: string;
  series_name: string;
  series_short_name: string;
  session_id: number;
  session_results: Array<Session>;
  special_event_type: number;
  start_time: string;
  subsession_id: number;
  track: Track;
  track_state: TrackState;
  weather: Weather;
};

export type Suit = {
  pattern: number;
  color1: string;
  color2: string;
  color3: string;
};

export type Track = {
  track_id: number;
  track_name: string;
  config_name: string;
  category_id: number | undefined;
  category: string | undefined;
  race_week_num: number | undefined;
};

export type TrackState = {
  leave_marbles: boolean;
  practice_rubber: number;
  qualify_rubber: number;
  warmup_rubber: number;
  race_rubber: number;
  practice_grip_compound: number;
  qualify_grip_compound: number;
  warmup_grip_compound: number;
  race_grip_compound: number;
};

export type UserIdCarClassIdsMapping = {
  userId: number;
  carClassIds: Array<number>;
};

export type UserIdYearsMapping = {
  userId: number;
  years: Array<number>;
};

export type Weather = {
  version: number;
  type: number;
  temp_units: number;
  temp_value: number;
  rel_humidity: number;
  fog: number;
  wind_dir: number;
  wind_units: number;
  wind_value: number;
  skies: number;
  weather_var_initial: number;
  weather_var_ongoing: number;
  allow_fog: boolean;
  track_water: number;
  precip_option: number;
  time_of_day: number;
  simulated_start_utc_time: string;
  simulated_start_utc_offset: number;
};
