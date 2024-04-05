import type { Track } from "$lib/types";

export default ({ tracks }: { tracks: Array<{ track: Track }> }) => {
  return tracks.reduce<Map<string, { trackName: string; configName: string }>>(
    (
      map: Map<string, { trackName: string; configName: string }>,
      trackObject
    ) => {
      const { track } = trackObject as { track: Track };
      const trackName = `${track.track_name} ${track.config_name}`;
      map.set(
        trackName
          .toLowerCase()
          .replace(/[\(\)\[\]]/g, "")
          .replace(/[ \/]/g, "-")
          .replace("---", "-")
          .replace("-n-a", ""),
        {
          trackName: track.track_name || "",
          configName: track.config_name || "",
        }
      );
      return map;
    },
    new Map([]) as Map<string, { trackName: string; configName: string }>
  );
};
