export type License = {
  min_license_level: number,
  license_group: number,
  group_name: string,
};

export type Session = {
  simsession_name: string,
  simsession_type_name: string,
  simsession_type: number,
  simsession_number: number,
  results: Array<Record<string,unknown>>
}
