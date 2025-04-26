import axios from "@/plugins/axios";

export interface CityWeatherInfo {
    city: string;
    temperature: number;
    weatherCondition: string;
}


const API_BASE_URL = 'http://localhost:8088';

export async function fetchWeather(city: string): Promise<CityWeatherInfo[]> {
    const encodedCity = encodeURIComponent(city.trim());

    try {
        const response = await axios.get(`/weather?city=${encodedCity}`);

        if (response.status !== 200) {
            console.error("Error fetching weather:", response.status, response.statusText, response.data);
            throw new Error(`HTTP error ${response.status}: ${response.statusText} - ${JSON.stringify(response.data)}`);
        }

        const data: CityWeatherInfo[] = response.data;
        return data;

    } catch (error) {
        console.error("Network or other error fetching weather:", error);
        throw error;
    }
}