package repo

import (
	"fmt"
	"log"
)

func ExampleMakePrList() {
	prList, err := MakeRepoQuery(Config{
		Trackers: []Tracker{
			{
				Name:       "git_appletini",
				Owner:      "darvoid",
				Identifier: "cenas",
			},
		},
		ReviewAmount:   10,
		PrAmount:       10,
		CommentsAmount: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prList.generatedQuery)
	// Output:
	// query PRs {
	//   babilon: repository(name: "ant-colony", owner: "god") {
	//     label(name: "bug") {
	//       pullRequests(
	//         orderBy: { field: CREATED_AT, direction: ASC }
	//         first: 10
	//         states: [OPEN]
	//       ) {
	//         ...getPullRequest
	//       }
	//     }
	//   }
	// }
	// fragment getPullRequest on PullRequestConnection {
	//   edges {
	//     node {
	//       id
	//       title
	//       url
	//       baseRefName
	//       headRefName
	//       reviewRequests {
	//         totalCount
	//       }
	//       reviewDecision
	//       createdAt
	//       permalink
	//       mergeable
	//       state
	//       ...getReviewer
	//     }
	//   }
	// }
	// fragment getReviewer on PullRequest {
	//   reviews(first: 10) {
	//     edges {
	//       node {
	//         state
	//         body
	//         comments(first: 10) {
	//           edges {
	//             node {
	//               body
	//             }
	//           }
	//         }
	//         author {
	//           login
	//         }
	//       }
	//     }
	//   }
	// }
}
