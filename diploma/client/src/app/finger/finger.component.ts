import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';
declare var $: any
import { FormBuilder, FormGroup, FormControl, Validators, FormArray } from '@angular/forms';

@Component({
  selector: 'app-finger',
  templateUrl: './finger.component.html',
  styleUrls: ['./finger.component.css']
})

export class FingerComponent implements OnInit {
  form: FormGroup;
  websiteList: any
  constructor(
    private apiService: ApiService,
    private formBuilder: FormBuilder
  ) {
    this.form = this.formBuilder.group({
      website: this.formBuilder.array([], [Validators.required])
    })
  }

  public students
  ngOnInit(): void {
    this.getStudents()
  }

  getStudents() {
    this.apiService.GetAllStudents().subscribe(res => {
      this.students = res.user;
      this.websiteList = this.students;
    });
  }

  handleFileInput($event) {
    let stud = this.submit().website
    console.log(stud, $event.target.files[0].type)

    this.apiService.UploadFile(stud,$event.target.files[0]).subscribe(res => {
      console.log(res)
    });
  }

  onCheckboxChange(e) {
    const website: FormArray = this.form.get('website') as FormArray;

    if (e.target.checked) {
      website.push(new FormControl(e.target.value));
    } else {
      const index = website.controls.findIndex(x => x.value === e.target.value);
      website.removeAt(index);
    }
  }

  submit() {
    return this.form.value;
  }
}
