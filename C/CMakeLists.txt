cmake_minimum_required(VERSION 3.21)
project(C)

set(CMAKE_CXX_STANDARD 14)

include_directories(.)
include_directories(day1)

add_executable(C
        day1/index.cpp day1/pdf.cpp)
