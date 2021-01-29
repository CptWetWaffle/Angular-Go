import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {HomeComponent} from './home/home.component';
import {FormsModule} from "@angular/forms";
import {HttpClientModule} from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ServiceWorkerModule } from '@angular/service-worker';
import { environment } from '../environments/environment'
import {AboutComponent} from "./about/about.component";
import {ExperienceComponent} from "./experience/experience.component";
import {InterestComponent} from "./interest/interest.component";
import {ProjectComponent} from "./project/project.component";
import {SkillComponent} from "./skill/skill.component";
import {VolunteeringComponent} from "./volunteering/volunteering.component";
import {EducationComponent} from "./education/education.component";
import {NavBarComponent} from "./nav-bar/nav-bar.component";

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    AboutComponent,
    EducationComponent,
    ExperienceComponent,
    InterestComponent,
    ProjectComponent,
    SkillComponent,
    VolunteeringComponent,
    NavBarComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ServiceWorkerModule.register('ngsw-worker.js', { enabled: environment.production })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
