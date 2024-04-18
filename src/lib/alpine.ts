import type { Alpine } from "alpinejs";
// @ts-expect-error
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
        if (index) {
          // @ts-expect-error
          this[id] = false;
        } else {
          // @ts-expect-error
          this[id] = !this[id];
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
