import { defineDb, defineTable, column } from "astro:db";

const Car = defineTable({
  columns: {
    ai_enabled: column.boolean(),
    allow_number_colors: column.boolean(),
    allow_sponsor1: column.boolean(),
    allow_sponsor2: column.boolean(),
    allow_wheel_color: column.boolean(),
    award_exempt: column.boolean(),
    car_dirpath: column.text(),
    car_id: column.number({ primaryKey: true }),
    car_name: column.text(),
    car_name_abbreviated: column.text(),
    car_types: column.json(),
    car_weight: column.number(),
    categories: column.json(),
    created: column.date(),
    first_sale: column.date(),
    forum_url: column.text({ optional: true }),
    free_with_subscription: column.boolean(),
    has_headlights: column.boolean(),
    has_multiple_dry_tire_types: column.boolean(),
    hp: column.number(),
    is_ps_purchasable: column.boolean(),
    max_power_adjust_pct: column.number(),
    max_weight_penalty_kg: column.number(),
    min_power_adjust_pct: column.number(),
    package_id: column.number(),
    pattern: column.number({ optional: true }),
    price: column.number(),
    price_display: column.text({ optional: true }),
    retired: column.boolean(),
    search_filters: column.text(),
    sku: column.number(),
    has_rain_capable_tire_types: column.boolean(),
    rain_enabled: column.boolean(),
  },
});

const CarClass = defineTable({
  columns: {
    car_class_id: column.number({ primaryKey: true }),
    cars_in_class: column.json(),
    cust_id: column.number(),
    name: column.text(),
    relative_speed: column.number(),
    short_name: column.text(),
    rain_enabled: column.boolean({ default: false }),
  },
});

const PastSeason = defineTable({
  columns: {
    active: column.boolean(),
    car_classes: column.json(),
    driver_changes: column.boolean(),
    fixed_setup: column.boolean(),
    license_group: column.number(),
    license_group_types: column.json(),
    official: column.boolean(),
    race_weeks: column.json(),
    season_id: column.number({ primaryKey: true }),
    season_name: column.text(),
    season_quarter: column.number(),
    season_short_name: column.text(),
    season_year: column.number(),
    series_id: column.number(),
    has_supersessions: column.boolean(),
  },
});

const Season = defineTable({
  columns: {
    active: column.boolean(),
    allowed_season_members: column.json({ optional: true }),
    car_class_ids: column.json(),
    car_switching: column.boolean(),
    car_types: column.json(),
    caution_laps_do_not_count: column.boolean(),
    complete: column.boolean(),
    cross_license: column.boolean(),
    driver_change_rule: column.number(),
    driver_changes: column.boolean(),
    drops: column.boolean(),
    enable_pitlane_collisions: column.boolean(),
    fixed_setup: column.boolean(),
    green_white_checkered_limit: column.number(),
    grid_by_class: column.boolean(),
    hardcore_level: column.number(),
    has_supersessions: column.boolean(),
    ignore_license_for_practice: column.boolean(),
    incident_limit: column.number(),
    incident_warn_mode: column.number(),
    incident_warn_param1: column.number(),
    incident_warn_param2: column.number(),
    is_heat_racing: column.boolean(),
    license_group: column.number(),
    license_group_types: column.json(),
    lucky_dog: column.boolean(),
    max_team_drivers: column.number(),
    max_weeks: column.number(),
    min_team_drivers: column.number(),
    multiclass: column.boolean(),
    must_use_diff_tire_types_in_race: column.boolean(),
    next_race_session: column.json({ optional: true }),
    num_opt_laps: column.number(),
    official: column.boolean(),
    op_duration: column.number(),
    open_practice_session_type_id: column.number(),
    qualifier_must_start_race: column.boolean(),
    race_week: column.number(),
    race_week_to_make_division: column.number({ optional: true }),
    reg_user_count: column.number(),
    region_competition: column.boolean(),
    restrict_by_member: column.boolean(),
    restrict_to_car: column.boolean(),
    restrict_viewing: column.boolean(),
    schedule_description: column.text(),
    schedules: column.json(),
    season_id: column.number({ primaryKey: true }),
    season_name: column.text(),
    season_quarter: column.number(),
    season_short_name: column.text(),
    season_year: column.number(),
    send_to_open_practice: column.boolean(),
    series_id: column.number(),
    short_parade_lap: column.boolean(),
    start_date: column.date(),
    start_on_qual_tire: column.boolean(),
    start_zone: column.boolean(),
    track_types: column.json(),
    unsport_conduct_rule_mode: column.number(),
  },
});

const Standing = defineTable({
  columns: {
    id: column.text({ primaryKey: true }),
    car_class_id: column.number({ references: () => CarClass.columns.car_class_id }),
    division: column.number(),
    overall_rank: column.number(),
    rank: column.number(),
    cust_id: column.number(),
    display_name: column.text(),
    club_id: column.number(),
    license: column.json(),
    helmet: column.json(),
    weeks_counted: column.number(),
    starts: column.number(),
    wins: column.number(),
    top5: column.number(),
    top25_percent: column.number(),
    poles: column.number(),
    avg_start_position: column.number(),
    avg_finish_position: column.number(),
    avg_field_size: column.number(),
    laps: column.number(),
    laps_led: column.number(),
    incidents: column.number(),
    points: column.number(),
    raw_points: column.number(),
    week_dropped: column.boolean(),
    country_code: column.text(),
    country: column.text(),
    season_id: column.number({ references: () => PastSeason.columns.season_id }),
    season_name: column.text(),
    division_rank: column.number({ default: 0 }),
  },
});

const Subsession = defineTable({
  columns: {
    allowed_licenses: column.json(),
    // associated_subsession_ids: column.json(),
    // can_protest: column.boolean(),
    car_classes: column.json(),
    // caution_type: column.number(),
    // cooldown_minutes: column.number(),
    corners_per_lap: column.number(),
    // damage_model: column.number(),
    // driver_change_param1: column.number(),
    // driver_change_param2: column.number(),
    // drive_change_rule: column.number(),
    // driver_changes: column.boolean(),
    end_time: column.date(),
    event_average_lap: column.number(),
    event_laps_complete: column.number(),
    event_strength_of_field: column.number(),
    // event_type: column.number(),
    // event_type_name: column.text(),
    // heat_info_id: column.number(),
    license_category: column.text(),
    license_category_id: column.number(),
    // limit_minutes: column.number(),
    // max_team_drivers: column.number(),
    max_weeks: column.number(),
    num_caution_laps: column.number(),
    num_cautions: column.number(),
    // num_laps_for_qual_average: column.number(),
    // num_laps_for_solo_average: column.number(),
    num_lead_changes: column.number(),
    // official_session: column.boolean(),
    // points_type: column.text(),
    // private_session_id: column.number(),
    race_summary: column.json(),
    race_week_num: column.number(),
    // results_restricted: column.boolean(),
    season_id: column.number({ references: () => PastSeason.columns.season_id }),
    season_name: column.text(),
    season_quarter: column.number(),
    season_short_name: column.text(),
    season_year: column.number(),
    series_id: column.number(),
    series_logo: column.text({ optional: true }),
    series_name: column.text(),
    series_short_name: column.text(),
    session_id: column.number(),
    // special_event_type: column.number(),
    start_time: column.date(),
    subsession_id: column.number({ primaryKey: true }),
    track: column.json(),
    // track_state: column.json(),
    // weather: column.json(),
  },
});

const SubsessionSessionResults = defineTable({
  columns: {
    id: column.text({ primaryKey: true }),
    subsession_id: column.number({ references: () => Subsession.columns.subsession_id }),
    simsession_number: column.number(),
    simsession_type: column.number(),
    simsession_type_name: column.text(),
    simsession_name: column.text(),
    cust_id: column.number(),
    display_name: column.text(),
    finish_position: column.number(),
    finish_position_in_class: column.number(),
    laps_lead: column.number(),
    laps_complete: column.number(),
    opt_laps_complete: column.number(),
    interval: column.number(),
    class_interval: column.number(),
    average_lap: column.number(),
    best_lap_num: column.number(),
    best_lap_time: column.number(),
    best_nlaps_num: column.number(),
    best_nlaps_time: column.number(),
    best_qual_lap_num: column.number(),
    best_qual_lap_time: column.number(),
    reason_out_id: column.number(),
    reason_out: column.number(),
    champ_points: column.number(),
    drop_race: column.boolean(),
    club_points: column.number({ default: 0 }),
    position: column.number(),
    qual_lap_time: column.number(),
    starting_position: column.number(),
    starting_position_in_class: column.number(),
    car_class_id: column.number(),
    car_class_name: column.text(),
    car_class_short_name: column.text(),
    club_id: column.number({ default: 0 }),
    club_name: column.text({ default: '' }),
    club_shortname: column.text({ default: '' }),
    division: column.number(),
    division_name: column.text({ optional: true }),
    old_license_level: column.number(),
    old_sub_level: column.number(),
    old_cpi: column.number(),
    oldi_rating: column.number(),
    old_ttrating: column.number(),
    new_license_level: column.number(),
    new_sub_level: column.number(),
    new_cpi: column.number(),
    newi_rating: column.number(),
    new_ttrating: column.number(),
    multiplier: column.number({ default: 0 }),
    license_change_oval: column.number(),
    license_change_road: column.number(),
    incidents: column.number(),
    max_pct_fuel_fill: column.number(),
    weight_penalty_kg: column.number(),
    league_points: column.number(),
    league_agg_points: column.number(),
    car_id: column.number(),
    car_name: column.text(),
    aggregate_champ_points: column.number(),
    livery: column.json(),
    suit: column.json(),
    helmet: column.json(),
    watched: column.boolean(),
    friend: column.boolean(),
    ai: column.boolean(),
  }
})

const User = defineTable({
  columns: {
    account: column.json(),
    ai: column.boolean({ default: false }),
    alpha_tester: column.boolean(),
    broadcaster: column.boolean(),
    bypass_hosted_password: column.boolean({ default: false }),
    car_packages: column.json(),
    club_id: column.number({ default: 0 }),
    club_name: column.text({ default: '' }),
    connection_type: column.text(),
    cust_id: column.number({ primaryKey: true }),
    dev: column.boolean(),
    display_name: column.text(),
    download_server: column.text(),
    email: column.text({ default: '' }),
    first_name: column.text(),
    flags: column.number(),
    flags_hex: column.text(),
    has_read_comp_rules: column.boolean(),
    has_read_pp: column.boolean(),
    has_read_tc: column.boolean(),
    helmet: column.json(),
    hundred_pct_club: column.boolean(),
    last_login: column.date(),
    last_name: column.text(),
    last_season: column.number(),
    last_test_car: column.number({ default: 0 }),
    last_test_track: column.number({ default: 0 }),
    licenses: column.json(),
    member_since: column.text(),
    on_car_name: column.text(),
    other_owned_packages: column.json(),
    race_official: column.boolean({ default: false }),
    read_comp_rules: column.date(),
    read_pp: column.date(),
    read_tc: column.date(),
    restrictions: column.json(),
    suit: column.json(),
    track_packages: column.json(),
    twenty_pct_discount: column.boolean(),
    username: column.text({ default: '' }),
    rain_test: column.boolean({ default: false }),
  },
});

// https://astro.build/db/config
export default defineDb({
  tables: {
    Car,
    CarClass,
    PastSeason,
    Season,
    Standing,
    Subsession,
    SubsessionPracticeResults: SubsessionSessionResults,
    SubsessionQualifyingResults: SubsessionSessionResults,
    SubsessionRaceResults: SubsessionSessionResults,
    User,
  },
});
