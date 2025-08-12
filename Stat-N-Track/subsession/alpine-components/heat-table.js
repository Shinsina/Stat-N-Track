const numberSort = ({ array, direction }) => {
  if (direction === "ascending") {
    return array.sort((a, b) => b - a);
  }
  return array.sort((a, b) => a - b);
};

const stringSort = ({ array, direction }) => {
  if (direction === "ascending") {
    return array.sort();
  }
  return array.sort().reverse();
};

export default function HEAT() {
  return {
    sort() {
      const [tableName, columnName, rowCount, direction] =
        this.$el.id.split(" ");
      const rows = new Map();
      const keys = new Set();
      for (let rowNumber = 0; rowNumber < Number(rowCount); rowNumber += 1) {
        const initialValue = this.$refs[
          `${tableName} ${columnName} ${rowNumber}`
        ].id
          .split("~")
          .pop();
        const value = Number.isNaN(Number(initialValue))
          ? initialValue
          : Number(initialValue);
        if (rows.get(value)) {
          rows.set(value, [...rows.get(value), rowNumber]);
        } else {
          rows.set(value, [rowNumber]);
          keys.add(value);
        }
      }
      const keysArray = Array.from(keys);
      const isNumbers = keysArray.filter(
        (key) => typeof key === "number"
      ).length;
      const internalSort = () => {
        if (isNumbers) {
          return numberSort({
            array: keysArray.map((value) => Number(value)),
            direction,
          });
        }
        return stringSort({
          array: keysArray.map((value) => String(value)),
          direction,
        });
      };
      const elementId = `${tableName} ${columnName} ${rowCount}`;
      this.$el.id =
        direction === "ascending" ? elementId : `${elementId} ascending`;
      const rowElements = internalSort()
        .map((key) => key)
        .map((key) => rows.get(key))
        .flat()
        .map((value) => this.$refs[`${tableName} ${value}`].outerHTML);
      const tableHeadings = this.$refs.Headings.outerHTML;
      this.$refs.Table.innerHTML = `${tableHeadings}${rowElements.join("")}`;
    },
  };
}
