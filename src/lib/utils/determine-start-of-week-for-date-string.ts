export default ({
  object,
  dateString,
}: {
  object: Record<string, Array<string>>;
  dateString: string;
}) => {
  const d = new Date(String(dateString));
  const day = d.getDay();
  const utcDateString = d.toUTCString();
  const minutesAwayFromUTC = d.getTimezoneOffset();
  const dayZero = minutesAwayFromUTC ? 1 : 0;
  if (day === dayZero && !object[utcDateString]) {
    object[utcDateString] = [utcDateString];
  } else if (day !== dayZero) {
    const distanceFromOne = day - dayZero;
    const dayOneOfWeek = new Date(
      Number(d) - distanceFromOne * 86400000
    ).toUTCString();
    if (object[dayOneOfWeek]) {
      object[dayOneOfWeek].push(utcDateString);
    } else {
      object[dayOneOfWeek] = [utcDateString];
    }
  }
  return object;
};
