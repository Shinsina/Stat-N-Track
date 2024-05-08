import type { Alpine } from "alpinejs";
import collapse from "@alpinejs/collapse";
import Table from "$lib/components/table";

export default (Alpine: Alpine) => {
  Alpine.plugin(collapse);
  Alpine.data("table", Table);
  Alpine.data("PRACTICE", Table);
  Alpine.data("QUALIFY", Table);
  Alpine.data("HEAT", Table);
  Alpine.data("RACE", Table);
  Alpine.data("tables", () => ({
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
  }));
  Alpine.data("dropdown", () => ({
    opened: false,
    toggle() {
      this.opened = !this.opened;
    },
  }));
};
