import { FetchWeather } from "../../wailsjs/go/main/App";

export interface CityWeatherInfo {
    city: string;
    temperature: number;
    weatherCondition: string;
}

export async function fetchWeather(city: string): Promise<CityWeatherInfo[]> {
    try {
        const data = await FetchWeather(city);
        return data;
    } catch (error: any) {
        console.error("Network or other error fetching weather:", error);
        throw new Error(error.message || "Failed to fetch weather data");
    }
}