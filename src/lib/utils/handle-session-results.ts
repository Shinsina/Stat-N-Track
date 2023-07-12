import parseLapTime from "./parse-lap-time";

type HandleResultsInput = { keysToDisplay: Set<String>, results: Array<Record<string, any>> };
type HandleResultsOuput = { keysArray: Array<string>, handledResults: Array<Record<string, any>> };

const zeroIndexedKey = new Set([
  'finish_position',
  'finish_position_in_class',
  'starting_position',
  'starting_position_in_class',
]);

const lapTimeField = new Set([
  'interval',
  'class_interval',
  'average_lap',
  'best_lap_time',
]);


export default ({ keysToDisplay, results }: HandleResultsInput): HandleResultsOuput => {
  const uniqueKeysArray: Set<string> = new Set([]);
  const handledResults = results
    .map((result) => Object.keys(result)
      .filter((key) => keysToDisplay.has(key))
        .map((key) => {
          uniqueKeysArray.add(key);
          if (zeroIndexedKey.has(key)) return result[key] + 1;
          if (lapTimeField.has(key)) return parseLapTime(result[key]);
          return result[key]
    }));
  const keysArray = Array.from(uniqueKeysArray)
  return { keysArray, handledResults };
}
