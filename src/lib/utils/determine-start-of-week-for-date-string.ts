export default ({
  object,
  dateString,
}: {
  object: Record<string, Array<string>>;
  dateString: string;
}) => {
  const d = new Date(`${String(dateString)} 19:00:00 GMT-05:00`);
  const day = d.getDay();
  const utcDateString = d.toUTCString();
  if (day === 1 && !object[utcDateString]) {
    object[utcDateString] = [utcDateString];
  } else if (day !== 1) {
    const distanceFromOne = day - 1;
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
