import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.css']
})
export class AddComponent implements OnInit {
  public addtype
  public date
  constructor(
    public activatedRoute: ActivatedRoute

  ) {
    this.activatedRoute.params.subscribe(async ({ child }) => {
      this.addtype = child
      this.date = new Date()
    })
  }

  ngOnInit(): void {

  }

}
