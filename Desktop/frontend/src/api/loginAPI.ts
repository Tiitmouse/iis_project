import {Login, Logout} from "../../wailsjs/go/api/Secure"

export async function login(username: string, password: string) {
  try {
   await Login(username,password)
  } catch (error: any) {
    console.error("Login error:", error);
    throw error;
  }
}

export async function logout() {
   try {
    await Logout()
   } catch (error: any) {
     console.error("Logout error: ", error);
     throw error;
   }
}