export interface Student {
  id: string; // 学生唯一标识（如学号）
  name: string; // 学生姓名
  age: number; // 年龄
  gender: "male" | "female" | "other"; // 性别，使用字面量类型限制选项
  email: string; // 邮箱
  grade: string; // 年级或班级（如 "Grade 10" 或 "Class A"）
  enrollmentDate: Date; // 注册日期
  status: "active" | "inactive" | "graduated"; // 学生状态
}
