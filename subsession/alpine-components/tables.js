export default function tables() {
  return {
    PRACTICE: false,
    QUALIFY: false,
    HEAT: false,
    RACE: false,
    toggle() {
      this.$el.id.split(" ").forEach((id, index) => {
        if (
          id === "PRACTICE" ||
          id === "QUALIFY" ||
          id === "HEAT" ||
          id === "RACE"
        ) {
          if (index) {
            this[id] = false;
          } else {
            this[id] = !this[id];
          }
        }
      });
    },
  }
}
