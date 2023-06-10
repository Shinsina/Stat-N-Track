export default (averageLap: number): number => {
	const ingestedAverageLap = String(averageLap);
	const decimal = String(ingestedAverageLap.match(/[0-9]{4}$/));
	const withDecimal = ingestedAverageLap.replace(decimal, `.${decimal}`);
	const [priorToDecimal, _] = withDecimal.split('.');
	if (priorToDecimal.length > 2) {
		const seconds = String(priorToDecimal.match(/[0-9]{2}$/));
		withDecimal.replace(seconds, `.${seconds}`);
	}
	return Number(withDecimal.replace(/0$/, '').trim());
}
