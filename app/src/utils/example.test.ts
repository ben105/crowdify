import { describe, it, expect } from "vitest";
import { isEvening } from "./example";

describe("isEvening", () => {
  it("returns false for morning time (6:00 AM)", () => {
    const morningTime = new Date();
    morningTime.setHours(6, 0, 0, 0);
    expect(isEvening(morningTime)).toBe(false);
  });

  it("returns false for noon (12:00 PM)", () => {
    const noonTime = new Date();
    noonTime.setHours(12, 0, 0, 0);
    expect(isEvening(noonTime)).toBe(false);
  });

  it("returns false for afternoon (3:30 PM)", () => {
    const afternoonTime = new Date();
    afternoonTime.setHours(15, 30, 0, 0);
    expect(isEvening(afternoonTime)).toBe(false);
  });

  it("returns false for early evening (5:59 PM)", () => {
    const earlyEveningTime = new Date();
    earlyEveningTime.setHours(17, 59, 0, 0);
    expect(isEvening(earlyEveningTime)).toBe(false);
  });

  it("returns true for evening start (6:00 PM)", () => {
    const eveningStartTime = new Date();
    eveningStartTime.setHours(18, 0, 0, 0);
    expect(isEvening(eveningStartTime)).toBe(true);
  });

  it("returns true for evening time (8:30 PM)", () => {
    const eveningTime = new Date();
    eveningTime.setHours(20, 30, 0, 0);
    expect(isEvening(eveningTime)).toBe(true);
  });

  it("returns true for late evening (11:59 PM)", () => {
    const lateEveningTime = new Date();
    lateEveningTime.setHours(23, 59, 59, 999);
    expect(isEvening(lateEveningTime)).toBe(true);
  });

  it("returns false for midnight (12:00 AM)", () => {
    const midnightTime = new Date();
    midnightTime.setHours(0, 0, 0, 0);
    expect(isEvening(midnightTime)).toBe(false);
  });

  it("handles date objects correctly", () => {
    const specificDate = new Date("2025-03-17T21:30:00"); // 9:30 PM
    expect(isEvening(specificDate)).toBe(true);
  });
});
