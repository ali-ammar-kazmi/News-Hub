import axios from "axios";
import { backend } from "../components/constants";

const getWeather = async (city) => {
  try {
    const response = await axios.get(`${backend}/weather?cityName=${city}`);
    return response; // Return the response data
  } catch (error) {
    console.error("Error fetching news:", error);
  }
};

export { getWeather };
