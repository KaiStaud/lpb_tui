cmake_minimum_required(VERSION 3.10)
project(CANOpen-TIF)

find_package(Curses REQUIRED)
#find_package(eCAL REQUIRED)
find_package(CLI11 CONFIG REQUIRED)
#find_package(Protobuf REQUIRED)

include_directories(${CURSES_INCLUDE_DIR} ${CMAKE_CURRENT_BINARY_DIR})

add_executable(${PROJECT_NAME} 
main.cpp
screens/manual.cpp
)
#set(drive_proto
#    ${CMAKE_CURRENT_SOURCE_DIR}/protobuf/drive.proto  
#)

#PROTOBUF_TARGET_CPP(${PROJECT_NAME} ${CMAKE_CURRENT_SOURCE_DIR}/protobuf ${drive_proto})

target_link_libraries(
        ${PROJECT_NAME} 
        ncurses 
        panel
        menu
	#       eCAL::core
	#protobuf::libprotobuf
        )
