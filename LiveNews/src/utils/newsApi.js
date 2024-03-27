import axios from "axios";
import { backend } from "../components/constants";

const getNews = async (category, page, pageSize) => {
  try {
    const response = await axios.get(`${backend}/news?category=${category}&page=${page}&pageSize=${pageSize}`);
    return response.data; // Return the response data
  } catch (error) {
    console.error("Error fetching news:", error);
  }
};

export { getNews };
