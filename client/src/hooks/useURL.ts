import { API } from "../api/api";
export default function useURL() {
  const CreateURL = async (data: {
    url: string;
    method: string[];
    expireIn: string;
  }) => {
    try {
      const res = await API.post("/api/url", data);
      return {
        success: true,
        ...res.data,
      };
    } catch (err: any) {
      return { success: false, message: err.response.data.message };
    }
  };
  return { CreateURL };
}
