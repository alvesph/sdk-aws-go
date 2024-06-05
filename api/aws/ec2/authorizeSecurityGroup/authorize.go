package authorizesecuritygroup

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func AuthorizeSG() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
	if err != nil {
		fmt.Println("Erro ao criar sessão:", err)
		return
	}
	svc := ec2.New(sess)

	groupID := "sg-0b8d419959042adc3"

	permission := &ec2.IpPermission{
		IpProtocol: aws.String("tcp"),
		FromPort:   aws.Int64(443),
		ToPort:     aws.Int64(443),
		IpRanges: []*ec2.IpRange{
			{
				CidrIp:      aws.String("0.0.0.0/0"),
				Description: aws.String("Allow HTTP traffic"),
			},
		},
	}

	authorizeRequest := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:       aws.String(groupID),
		IpPermissions: []*ec2.IpPermission{permission},
	}

	_, err = svc.AuthorizeSecurityGroupIngress(authorizeRequest)
	if err != nil {
		fmt.Println("Erro ao autorizar a regra de segurança:", err)
		return
	}

	fmt.Println("Regra de segurança criada com sucesso!")
}
