package timeout

import "context"

type AcceptanceFunc func(string) (string, error)

type PopulatedFunc func(context.Context, string) (string, error)

func Timeout(f AcceptanceFunc) PopulatedFunc {
	// returns populated function
	return func(ctx context.Context, arg string) (string, error) {

		// make channel by the type of acceptance function return
		resCh := make(chan string)
		errCh := make(chan error)

		// set a goroutine to run the acceptance function and wait it till ends
		go func() {
			res, err := f(arg)
			resCh <- res
			errCh <- err
		}()

		// we wait to the response till context times out
		select {
		case res := <-resCh:
			return res, <-errCh
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}
