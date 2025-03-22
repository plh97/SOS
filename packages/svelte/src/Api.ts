import axios, { type AxiosRequestConfig } from "axios";
import type { Student } from "./interface";

const instance = axios.create({
  //   withCredentials: true,
  baseURL: "/api",
});

// Add a request interceptor
instance.interceptors.request.use(
  function (config) {
    // Do something before request is sent
    return config;
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error);
  }
);

// Add a response interceptor
instance.interceptors.response.use(
  function (response) {
    return response.data;
  },
  function (error) {
    return Promise.reject(error);
  }
);

export default function request<T = unknown>(option: AxiosRequestConfig) {
  return new Promise<T>((resolve, reject) => {
    instance
      .request<T>(option)
      .then((res) => resolve(res as T))
      .catch((err) => reject(err));
  });
}

export const getStudentList = async (): Promise<Student[]> => {
  const res: any = await request({
    url: "/v1/student",
  });

  return res;
};

export const createStudent = async (data: Student): Promise<Student> => {
  return request({
    method: "post",
    url: "/v1/student",
    data,
  });
};

export const deleteStudent = async (id: string): Promise<Student> => {
  return request({
    method: "delete",
    url: `/v1/student/${id}`,
  });
};
