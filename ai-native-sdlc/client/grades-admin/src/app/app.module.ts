import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http'; // added
import { AppComponent } from './app.component';
import { GradesComponent } from './grades/grades.component';

@NgModule({
  declarations: [
    AppComponent,
    GradesComponent
    // ...existing declarations...
  ],
  imports: [
    BrowserModule,
    HttpClientModule, // added
    // ...existing imports...
  ],
  providers: [
    // ...existing providers...
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }