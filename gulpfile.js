var gulp = require('gulp');
var sass = require('gulp-sass');
var rename = require("gulp-rename");

gulp.task('styles', function() {
  return gulp.src('./app/*/assets/css/sass/**/*.scss', {base: './'})
    .pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
    .pipe(rename(function(path){
      path.dirname = path.dirname.replace('sass', '')
    }))
    .pipe(gulp.dest('.'));
});

//Watch task
gulp.task('default', function() {
  return gulp.watch('./app/*/assets/css/sass/**/*.scss', gulp.series('styles'));
});
