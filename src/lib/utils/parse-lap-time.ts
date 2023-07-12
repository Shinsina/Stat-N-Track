export default (lap: number): string => {
  const ingestedLap = String(lap);
  const placeDecimal = () => {
    switch (ingestedLap.length) {
      case 1:
        return `0.000${ingestedLap}`;
      case 2:
        return `0.00${ingestedLap}`;
      case 3:
        return `0.0${ingestedLap}`;
      case 4:
        return `0.${ingestedLap}`;
      default:
        const decimal = ingestedLap.match(/[0-9]{4}$/);
        return ingestedLap.replace(String(decimal), `.${decimal}`);
    }
  };
  const withDecimal = lap > 0 ? placeDecimal() : ingestedLap;
  const [priorToDecimal, _] = withDecimal.split(".");
  if (priorToDecimal.length > 2) {
    const seconds = String(priorToDecimal.match(/[0-9]{2}$/));
    withDecimal.replace(seconds, `.${seconds}`);
  }
  const numericalValue = Number(withDecimal.replace(/0$/, "").trim());
  // If the lap is a minute or longer
  if (numericalValue >= 60) {
    const minutes = Math.floor(numericalValue / 60);
    const seconds = Math.floor(numericalValue - minutes * 60);
    const subseconds = withDecimal.split(".").pop();
    const secondsString = seconds >= 10 ? seconds : `0${seconds}`;
    return `${minutes}:${secondsString}.${subseconds}`;
  }
  return withDecimal.replace(/0$/, "").trim();
};
