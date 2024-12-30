// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of the log anomaly detectors in the account.
func (c *Client) ListLogAnomalyDetectors(ctx context.Context, params *ListLogAnomalyDetectorsInput, optFns ...func(*Options)) (*ListLogAnomalyDetectorsOutput, error) {
	if params == nil {
		params = &ListLogAnomalyDetectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLogAnomalyDetectors", params, optFns, c.addOperationListLogAnomalyDetectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLogAnomalyDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLogAnomalyDetectorsInput struct {

	// Use this to optionally filter the results to only include anomaly detectors
	// that are associated with the specified log group.
	FilterLogGroupArn *string

	// The maximum number of items to return. If you don't specify a value, the
	// default maximum value of 50 items is used.
	Limit *int32

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLogAnomalyDetectorsOutput struct {

	// An array of structures, where each structure in the array contains information
	// about one anomaly detector.
	AnomalyDetectors []types.AnomalyDetector

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLogAnomalyDetectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLogAnomalyDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLogAnomalyDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLogAnomalyDetectors"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLogAnomalyDetectors(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// ListLogAnomalyDetectorsPaginatorOptions is the paginator options for
// ListLogAnomalyDetectors
type ListLogAnomalyDetectorsPaginatorOptions struct {
	// The maximum number of items to return. If you don't specify a value, the
	// default maximum value of 50 items is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLogAnomalyDetectorsPaginator is a paginator for ListLogAnomalyDetectors
type ListLogAnomalyDetectorsPaginator struct {
	options   ListLogAnomalyDetectorsPaginatorOptions
	client    ListLogAnomalyDetectorsAPIClient
	params    *ListLogAnomalyDetectorsInput
	nextToken *string
	firstPage bool
}

// NewListLogAnomalyDetectorsPaginator returns a new
// ListLogAnomalyDetectorsPaginator
func NewListLogAnomalyDetectorsPaginator(client ListLogAnomalyDetectorsAPIClient, params *ListLogAnomalyDetectorsInput, optFns ...func(*ListLogAnomalyDetectorsPaginatorOptions)) *ListLogAnomalyDetectorsPaginator {
	if params == nil {
		params = &ListLogAnomalyDetectorsInput{}
	}

	options := ListLogAnomalyDetectorsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLogAnomalyDetectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLogAnomalyDetectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLogAnomalyDetectors page.
func (p *ListLogAnomalyDetectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLogAnomalyDetectorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListLogAnomalyDetectors(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListLogAnomalyDetectorsAPIClient is a client that implements the
// ListLogAnomalyDetectors operation.
type ListLogAnomalyDetectorsAPIClient interface {
	ListLogAnomalyDetectors(context.Context, *ListLogAnomalyDetectorsInput, ...func(*Options)) (*ListLogAnomalyDetectorsOutput, error)
}

var _ ListLogAnomalyDetectorsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListLogAnomalyDetectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLogAnomalyDetectors",
	}
}