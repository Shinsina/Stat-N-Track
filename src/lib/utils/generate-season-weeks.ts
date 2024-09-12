import determineStartOfWeekForDateString from "$lib/utils/determine-start-of-week-for-date-string";

export default ({
  seasons,
}: {
  seasons: Array<{ schedules: Array<{ start_date: string }> }>;
}) => {
  const seasonStart = Date.UTC(2024, 8, 10, 0, 0, 0, 0);
  const seasonEnd = Date.UTC(2024, 11, 4, 0, 0, 0, 0);
  const startDates = seasons.reduce(
    (
      set: Set<string>,
      season: { schedules: Array<{ start_date: string }> }
    ) => {
      const { schedules } = season;
      if (Array.isArray(schedules)) {
        schedules.forEach((schedule) => {
          const { start_date } = schedule;
          const [year, month, day] = start_date.split("-");
          const date = Date.UTC(
            Number(year),
            Number(month) - 1,
            Number(day),
            0,
            0,
            0,
            0
          );
          if (date >= seasonStart && date <= seasonEnd) {
            set.add(start_date);
          }
        });
      }
      return set;
    },
    new Set([]) as Set<string>
  );
  const weeks = Array.from(startDates)
    .sort()
    .reduce<Record<string, Array<string>>>(
      (object, dateString) =>
        determineStartOfWeekForDateString({ object, dateString }),
      {}
    );
  return weeks;
};
