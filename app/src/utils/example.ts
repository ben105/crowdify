export function isEvening(date: Date = new Date()): boolean {
    const hours = date.getHours();
    return hours >= 18;
}