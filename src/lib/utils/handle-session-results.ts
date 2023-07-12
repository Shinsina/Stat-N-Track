type HandleResultsInput = {
  keysToDisplay: Set<String>;
  results: Array<Record<string, any>>;
};
type HandleResultsOuput = {
  keysArray: Array<string>;
  handledResults: Array<Record<string, any>>;
};

const zeroIndexedKey = new Set([
  "finish_position",
  "finish_position_in_class",
  "starting_position",
  "starting_position_in_class",
]);

export default ({
  keysToDisplay,
  results,
}: HandleResultsInput): HandleResultsOuput => {
  const uniqueKeys: Set<string> = new Set([]);
  const handledResults = results.map((result) =>
    Object.keys(result)
      .filter((key) => keysToDisplay.has(key))
      .map((key) => {
        uniqueKeys.add(key);
        if (zeroIndexedKey.has(key)) return result[key] + 1;
        return result[key];
      })
  );
  const keysArray = Array.from(uniqueKeys);
  return { keysArray, handledResults };
};
