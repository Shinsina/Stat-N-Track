export type Session = {
  simsession_name: string,
  simsession_type_name: string,
  simsession_type: number,
  simsession_number: number,
  results: Array<Record<string,unknown>>
}
