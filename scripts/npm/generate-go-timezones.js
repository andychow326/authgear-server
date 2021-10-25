const data = require("tzdata");
const { IANAZone } = require("luxon");

function getTimezoneNames() {
  const names = [];

  for (const [key, value] of Object.entries(data.zones)) {
    // This is an alias.
    if (typeof value === "string") {
      continue;
    }
    if (!key.includes("/")) {
      continue;
    }
    if (key.startsWith("Etc/")) {
      continue;
    }

    const iana = IANAZone.create(key);
    if (!iana.isValid) {
      continue;
    }

    names.push(key);
  }

  return names;
}

function main() {
  const f = console.log.bind(console);
  const names = getTimezoneNames();

  f("package tzutil");
  f();
  f("// This file is generated by generate-go-timezones.js");
  f("// Do NOT edit it manually!");
  f();
  f("var timezoneNames = []string{");
  for (const name of names) {
    f('\t"' + name + '",');
  }
  f("}");
}

main();