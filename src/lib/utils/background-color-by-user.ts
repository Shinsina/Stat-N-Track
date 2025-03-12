const userIdMap: Map<number, string> = new Map([
  [182407, "bg-amber-700"],
  [188056, "bg-gray-700"],
  [251134, "bg-red-700"],
  [300752, "bg-orange-700"],
  [331322, "bg-indigo-700"],
  [589449, "bg-blue-700"],
  [714312, "bg-yellow-700"],
  [746377, "bg-green-700"],
  [815162, "bg-purple-700"],
  [908575, "bg-emerald-700"]
]);

const displayNameMap: Map<string, string> = new Map([
  ["Antonio Estrada", "bg-amber-700"],
  ["Doug South", "bg-gray-700"],
  ["Kyle Klendworth", "bg-red-700"],
  ["Jacob Collins", "bg-orange-700"],
  ["Jesper Öhrman", "bg-indigo-700"],
  ["Bryan Campbell2", "bg-blue-700"],
  ["Sam Karasala", "bg-yellow-700"],
  ["Ty Quila", "bg-green-700"],
  ["Jack Glenzinski", "bg-purple-700"],
  ["Cody Cavaco", "bg-emerald-700"]
]);

export const bgColorByUserId = (userId: number): string =>
  userIdMap.get(userId) || "";

export const bgColorByUserDisplayName = (displayName: string): string =>
  displayNameMap.get(displayName) || "";
