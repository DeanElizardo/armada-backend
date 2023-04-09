import { prisma } from "./index";
import { Course } from "@prisma/client";

interface ICourseDetails {
  name: string;
}

// Create a Course //!REFACTORED
async function createCourse(name: string, cohortId: number) {
  const course = await prisma.course.create({
    data: {
      name,
      cohortId,
    },
  });

  return course;
}

// Delete a Course //!REFACTORED
export async function deleteCourse(id: number) {
  const course = await prisma.course.delete({
    where: { id },
  });

  return course;
}

// Get All Courses //!REFACTORED
export async function retrieveAllCourses() {
  const courses = await prisma.course.findMany({
    include: {
      cohort: true,
    },
  });
  return courses;
}

// Get All Courses from a Specific Cohort //!REFACTORED
export async function retrieveAllCoursesFromCohort(cohortId: number) {
  const courses = await prisma.course.findMany({
    where: { cohortId },
  });

  return courses;
}

// Get All Users from a Specific Course //!REFACTORED
export async function retrieveAllUsersFromSpecificCourse(courseId: number) {
  const users = await prisma.user_Course.findMany({
    where: { courseId },
    include: {
      user: true,
    },
  });

  return users;
}

// Get Courses from List //! NEVER USED; DO NOT REFACTOR
export async function retrieveCoursesFromList(list: number[]) {
  const cohorts = await prisma.course.findMany({
    where: {
      id: { in: list },
    },
  });

  return cohorts;
}

// Get a Specific Course //!REFACTORED
export async function retrieveSpecificCourse(id: number) {
  const course = await prisma.course.findUnique({
    where: { id },
  });

  return course;
}

// Update a Course //!REFACTORED
export async function updateCourse(id: number, courseDetails: ICourseDetails) {
  const updateCourse = await prisma.course.update({
    where: { id },
    data: {
      ...courseDetails,
    },
  });

  return updateCourse;
}

const courseActions = {
  createCourse,
  deleteCourse,
  retrieveAllCourses,
  retrieveAllCoursesFromCohort,
  retrieveCoursesFromList,
  retrieveAllUsersFromSpecificCourse,
  retrieveSpecificCourse,
  updateCourse,
};

export default courseActions;
