import axios from "axios"
import { backendUrl, Header } from "../utils/urls"
import { LoginData, RegisterData, ChangePasswordPayload } from "../utils/types";

const url = "users/"


export const loginRequest = (data: LoginData) => {
    return axios.post(backendUrl + url + "login", data, Header())  // POST /users/login - bez JWT
}


export const registerRequest = (data: RegisterData) => {
    return axios.post(backendUrl + url, data, Header())  // POST /users - rejestracja wymaga JWT
}

export const changePassword = (data: ChangePasswordPayload) => {
    return axios.post(backendUrl + url + "change-password", data, Header())  // POST /users/change-password
}

export const getUsersData = () => {
    return axios.get(backendUrl + url, Header())  // GET /users, JWT + admin role
}

export const deleteUserByEmail = (email: String) => {
    return axios.delete(backendUrl + url + email, Header())  // DELETE /users/:email, JWT + admin role
}

export const changePasswordForAdmin = (email: String, data: ChangePasswordPayload) => {
    return axios.post(backendUrl + url + `change-password/${email}`, data, Header())  // POST /users/change-password/:email, JWT + admin role
}