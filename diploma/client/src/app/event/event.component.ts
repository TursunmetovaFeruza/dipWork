import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-event',
  templateUrl: './event.component.html',
  styleUrls: ['./event.component.css']
})
export class EventComponent implements OnInit {

  constructor(
    private apiService: ApiService,

  ) { }

  ngOnInit(): void {
    this.getTables()
  }


  public lection
  public student
  public master
  public startTime
  public endTime
  // public faculties
  // public schools
  public masters
  // public courses
  // public groups
  // public langs
  public lections
  public students
  getTables() {
    this.apiService.GetAllStudents().subscribe(res => {
      this.students = res.user
    });
    // this.apiService.GetFaculties().subscribe(res => {
    //   this.faculties = res.user
    // });
    // this.apiService.GetSchools().subscribe(res => {
    //   this.schools = res.user
    // });
    this.apiService.GetAllMaters().subscribe(res => {
      this.masters = res.user
    });
    // this.apiService.GetCourses().subscribe(res => {
    //   this.courses = res.user
    // });
    // this.apiService.GetLanguages().subscribe(res => {
    //   this.langs = res.user
    // });
    this.apiService.GetLections().subscribe(res => {
      this.lections = res.user
    });
  }

  getEventValue() {
    console.log(this.lection, this.student, this.master, this.startTime, this.endTime)
    let event = {
      id: 0,
      lectionid: this.lection,
      masterid: this.master,
      startTime: getTime(this.startTime),
      endTime: getTime(this.endTime),
      studentid: this.student
    }
   
    function getTime(time) {
      return {
        "year": new Date(time).getFullYear(),
        "month": new Date(time).getMonth(),
        "day": new Date(time).getDate(),
        "hour": new Date(time).getHours(),
        "minute": new Date(time).getMinutes(),
      }
    }
    // new Date().toISOString()
    this.apiService.SetAttendance(event).subscribe(res => {
      console.log(res)
    })
  }
}
