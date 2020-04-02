import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { BarComponent } from './bar/bar.component';
import { LoginComponent } from './login/login.component'
import { StudentMainComponent } from './student-main/student-main.component';

const routes: Routes = [
  {
    path: 'home',
    component: HomeComponent
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'bar',
    component: BarComponent,
    children: [
      {
        path: 'student-main',
        component: StudentMainComponent
      },
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
