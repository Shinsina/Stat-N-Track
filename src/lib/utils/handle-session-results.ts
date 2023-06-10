import parseLapTime from "./parse-lap-time";

type HandleSessionResultsInput = { keysToDisplay: Set<String>, results: Array<Record<string, any>> };

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


export default ({ keysToDisplay, results }: HandleSessionResultsInput): Array<Record<string, any>> => {
  return results
    .map((result) => Object.keys(result)
      .filter((key) => keysToDisplay.has(key))
        .map((key) => {
          if (zeroIndexedKey.has(key)) return result[key] + 1;
          if (lapTimeField.has(key)) return parseLapTime(result[key]);
          return result[key]
    }));
}
