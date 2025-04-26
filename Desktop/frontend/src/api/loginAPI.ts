import axios from "@/plugins/axios";

interface LoginResponse {
  access_token: string;
  refresh_token: string;
}

export async function login(username: string, password: string): Promise<LoginResponse> {
  try {
    const response = await axios.post<LoginResponse>('/api/login', {
      username: username,
      password: password
    });

    if (response.status === 200) {
      return response.data;
    } else {
      throw new Error(`Login failed with status: ${response.status}`);
    }
  } catch (error: any) {
    console.error("Login error:", error);
    throw error;
  }
}