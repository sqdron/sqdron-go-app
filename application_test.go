package sqapp

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestApplicationMainflow(t *testing.T) {
	Convey("Create application", t, func() {
		app := App();
		So(app, ShouldNotBeNil)
		Convey("Use some middleware", func() {
			appWithMiddlewares := app.Use("m1").Use("m2");
			Convey("Run application", func() {
				appWithMiddlewares.Run();
			})
		});


		//user := &User{};
		//user.ID = "1";
		//user.Username = "Denis"
		//modelQuery := qr.For(user);
		//Convey("Make where query", func() {
		//	filterQuery := modelQuery.Where();
		//	Convey("Make Select", func() {
		//		selectQuery := filterQuery.Select();
		//		Convey("Make Order", func() {
		//			orderQuery := selectQuery.Order();
		//			Convey("Make ALL result", func() {
		//				result := orderQuery.All();
		//				So(result, ShouldNotBeNil)
		//			})
		//
		//			Convey("Make First result", func() {
		//				result := orderQuery.First();
		//				So(result, ShouldNotBeNil)
		//			})
		//
		//			Convey("Make Single result", func() {
		//				result := orderQuery.Single();
		//				So(result, ShouldNotBeNil)
		//			})
		//		})
		//	})
		//})
	})
}
