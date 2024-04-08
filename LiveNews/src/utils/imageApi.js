import axios from "axios";
import { backend } from "../components/constants";

const getImage = async (title) => {
  try {
    const response = await axios.get(`${backend}/image?title=${title}`);
    return response; // Return the response data
  } catch (error) {
    console.error("Error fetching news:", error);
  }
};

export { getImage };
