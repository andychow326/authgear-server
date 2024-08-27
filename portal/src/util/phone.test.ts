/* global it, expect */
import {
  cleanRawInputValue,
  trimCountryCallingCode,
  makePartialValue,
} from "./phone";

import { default as parsePhoneNumber } from "libphonenumber-js";

it("cleanRawInputValue", () => {
  expect(cleanRawInputValue("1234 1234")).toEqual("12341234");
  expect(cleanRawInputValue("asdf")).toEqual("");
  expect(cleanRawInputValue("我")).toEqual("");
  expect(cleanRawInputValue("+852+852")).toEqual("+852852");
});

it("trimCountryCallingCode", () => {
  expect(trimCountryCallingCode("+", "852")).toEqual("");
  expect(trimCountryCallingCode("+852", "852")).toEqual("");
  expect(trimCountryCallingCode("+85298", "852")).toEqual("98");
  expect(trimCountryCallingCode("98765432", "852")).toEqual("98765432");
});

it("makePartialValue", () => {
  expect(makePartialValue("+", "852")).toEqual("+852");
  expect(makePartialValue("+852", "852")).toEqual("+852");
  expect(makePartialValue("123", "852")).toEqual("+852123");
});

it("libphonenumber-js parsePhoneNumber", () => {
  // invalid
  expect(parsePhoneNumber("")).toEqual(undefined); // need digits
  expect(parsePhoneNumber("+")).toEqual(undefined); // need digits
  expect(parsePhoneNumber("+  ")).toEqual(undefined); // need digits
  expect(parsePhoneNumber("+-")).toEqual(undefined); // need digits
  expect(parsePhoneNumber("+852")).toEqual(undefined); // need number
  // valid
  expect(parsePhoneNumber("+852 9 9 9  99 999")).toHaveProperty(
    "number",
    "+85299999999"
  ); // allow space in middle

  expect(parsePhoneNumber("   +85299999999     ")).toHaveProperty(
    "number",
    "+85299999999"
  ); // trimmed
  expect(parsePhoneNumber("\t\r\n+85299999999\t\r\n")).toHaveProperty(
    "number",
    "+85299999999"
  ); // trimmed
  expect(parsePhoneNumber("+852-9999-9999")).toHaveProperty(
    "number",
    "+85299999999"
  ); // allow dashes in middle
});
